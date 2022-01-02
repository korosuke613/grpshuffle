package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1" // here
	"google.golang.org/grpc/keepalive"
)

func Connect(host string, port int) (conn *grpc.ClientConn, err error) {
	// see https://pkg.go.dev/google.golang.org/grpc/keepalive#ClientParameters
	kp := keepalive.ClientParameters{
		Time: 60 * time.Second,
	}

	addr := fmt.Sprintf("%v:%v", host, port)

	// insecure.NewCredentials() を指定することで、TLS ではなく平文で接続
	// 通信内容が保護できないし、不正なサーバーに接続しても検出できないので本当はダメ
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kp))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseConnect(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("close failed for healthHandler()"))
		return
	}
}

func Shuffle(conn *grpc.ClientConn, partition int32, targets []string) ([]*grpshuffle.Combination, error) {
	cc := grpshuffle.NewComputeClient(conn)

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
		return nil, err
	}

	return res.Combinations, nil
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(conn *grpc.ClientConn) (*HealthCheckResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel func()) {
		time.Sleep(2500 * time.Millisecond)
		cancel()
	}(cancel)

	healthClient := health.NewHealthClient(conn)

	res, err := healthClient.Check(ctx, &health.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	result := &HealthCheckResponse{
		Status: health.HealthCheckResponse_ServingStatus_name[int32(res.Status)],
	}

	return result, nil
}
