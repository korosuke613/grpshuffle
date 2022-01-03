package grpshuffle_client_test

import (
	"context"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	health "google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
)

type mockComputeClient struct {
	grpshuffle.ComputeClient
	mockShuffle func(ctx context.Context, in *grpshuffle.ShuffleRequest) (grpshuffle.ShuffleResponse, error)
}

func (m *mockComputeClient) Shuffle(ctx context.Context, in *grpshuffle.ShuffleRequest, opts ...grpc.CallOption) (*grpshuffle.ShuffleResponse, error) {
	result, err := m.mockShuffle(ctx, in)
	return &result, err
}

type mockHealthClient struct {
	health.HealthClient
	mockCheck func(ctx context.Context, in *health.HealthCheckRequest) (health.HealthCheckResponse, error)
}

func (m *mockHealthClient) Check(ctx context.Context, in *health.HealthCheckRequest, opts ...grpc.CallOption) (*health.HealthCheckResponse, error) {
	result, err := m.mockCheck(ctx, in)
	return &result, err
}
