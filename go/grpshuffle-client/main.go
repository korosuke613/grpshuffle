package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1" // here
	"google.golang.org/grpc/keepalive"
)

func main() {
	err := subMain()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func subMain() error {
	addr := os.Args[1]

	// see https://pkg.go.dev/google.golang.org/grpc/keepalive#ClientParameters
	kp := keepalive.ClientParameters{
		Time: 60 * time.Second,
	}

	// insecure.NewCredentials() を指定することで、TLS ではなく平文で接続
	// 通信内容が保護できないし、不正なサーバーに接続しても検出できないので本当はダメ
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kp))
	if err != nil {
		return err
	}

	defer conn.Close()

	cc := grpshuffle.NewComputeClient(conn)

	method := string(os.Args[2])
	if err != nil {
		return err
	}
	switch method {
	case "shuffle":
		{
			if len(os.Args) <= 5 {
				return errors.New("usage: client HOST:PORT METHOD PARTITION TARGET_1 TARGET_2 ... TARGET_N\nMETHOD: shuffle, health")
			}

			p, err := strconv.Atoi(os.Args[3])
			if err != nil {
				return err
			}
			s := os.Args[4:]

			err = callShuffle(cc, int32(p), s)
			if err != nil {
				return err
			}
		}
	case "health":
		{
			err := callHealth(conn)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("usage: client HOST:PORT METHOD PARTITION TARGET_1 TARGET_2 ... TARGET_N\nMETHOD: shuffle, health")
	}
	return nil
}

func callShuffle(cc grpshuffle.ComputeClient, partition int32, targets []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel func()) {
		time.Sleep(2500 * time.Millisecond)
		cancel()
	}(cancel)

	res, err := cc.Shuffle(ctx, &grpshuffle.ShuffleRequest{
		Targets:   targets,
		Partition: partition,
	})
	if err != nil {
		return err
	}

	result, err := json.MarshalIndent(res.Combinations, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(result))

	return nil
}

func callHealth(conn *grpc.ClientConn) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel func()) {
		time.Sleep(2500 * time.Millisecond)
		cancel()
	}(cancel)

	healthClient := health.NewHealthClient(conn)

	res, err := healthClient.Check(ctx, &health.HealthCheckRequest{})
	if err != nil {
		return err
	}

	status := map[string]string{
		"status": health.HealthCheckResponse_ServingStatus_name[int32(res.Status)],
	}
	result, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(result))

	return nil
}
