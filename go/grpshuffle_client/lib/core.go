package grpshuffle_client

import (
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1" // here
	"google.golang.org/grpc/keepalive"
)

func makeDialOpts(noTls bool) []grpc.DialOption {
	// see https://pkg.go.dev/google.golang.org/grpc/keepalive#ClientParameters
	kp := keepalive.ClientParameters{
		Time: 60 * time.Second,
	}

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithKeepaliveParams(kp))

	if noTls {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsConfig := tls.Config{
			InsecureSkipVerify: true,
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
	}

	return dialOpts
}

// Connect is create grpc.ClientConn
func Connect(host string, port int, noTls bool) (conn *grpc.ClientConn, err error) {
	addr := fmt.Sprintf("%v:%v", host, port)

	dialOpts := makeDialOpts(noTls)

	conn, err = grpc.NewClient(addr, dialOpts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// CloseConnect is close grpc.ClientConn
func CloseConnect(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("close failed for healthHandler()"))
		return
	}
}

// Shuffle is request to grpshuffle.ComputeClient
func Shuffle(cc *grpshuffle.ComputeClient, divide uint64, targets []string) ([]*grpshuffle.Combination, error) {

	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel func()) {
		time.Sleep(2500 * time.Millisecond)
		cancel()
	}(cancel)

	res, err := (*cc).Shuffle(ctx, &grpshuffle.ShuffleRequest{
		Targets: targets,
		Divide:  divide,
	})
	if err != nil {
		return nil, err
	}

	return res.Combinations, nil
}

// HealthCheckResponse is response HealthCheck
type HealthCheckResponse struct {
	Status string `json:"status"`
}

// HealthCheck is request to grpc_health_v1.HealthClient
func HealthCheck(hc *health.HealthClient) (*HealthCheckResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel func()) {
		time.Sleep(2500 * time.Millisecond)
		cancel()
	}(cancel)

	res, err := (*hc).Check(ctx, &health.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}

	result := &HealthCheckResponse{
		Status: health.HealthCheckResponse_ServingStatus_name[int32(res.Status)],
	}

	return result, nil
}
