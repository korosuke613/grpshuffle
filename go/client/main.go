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
	if len(os.Args) <= 4 {
		return errors.New("usage: client HOST:PORT PARTITION TARGET_1 TARGET_2 ... TARGET_N")
	}
	addr := os.Args[1]

	// see https://pkg.go.dev/google.golang.org/grpc/keepalive#ClientParameters
	kp := keepalive.ClientParameters{
		Time: 60 * time.Second,
	}

	// grpc.WithInsecure() を指定することで、TLS ではなく平文で接続
	// 通信内容が保護できないし、不正なサーバーに接続しても検出できないので本当はダメ
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithKeepaliveParams(kp))
	if err != nil {
		return err
	}

	defer conn.Close()

	cc := grpshuffle.NewComputeClient(conn)

	p, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return err
	}
	s := os.Args[3:]

	err = callShuffle(cc, int32(p), s)
	if err != nil {
		return err
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
