package grpshuffle_client_test

import (
	"context"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	grpshuffleClient "github.com/korosuke613/grpshuffle/go/grpshuffle_client/lib"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	type args struct {
		cc grpshuffle.ComputeClient
		d  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    []*grpshuffle.Combination
		wantErr bool
	}{
		{
			name: "Return Combination",
			args: args{
				cc: &mockComputeClient{
					mockShuffle: func(ctx context.Context, in *grpshuffle.ShuffleRequest) (grpshuffle.ShuffleResponse, error) {
						return grpshuffle.ShuffleResponse{
							Combinations: []*grpshuffle.Combination{
								{
									Targets: []string{"a", "b"},
								},
								{
									Targets: []string{"c", "d"},
								},
							},
						}, nil
					},
				},
				d: 3 * time.Second,
			},
			want: []*grpshuffle.Combination{
				{
					Targets: []string{"a", "b"},
				},
				{
					Targets: []string{"c", "d"},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := grpshuffleClient.Shuffle(&tt.args.cc, 2, []string{"a", "b", "c", "d"})
			if (err != nil) != tt.wantErr {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("actual %s, want %s", actual, tt.want)
			}

		})
	}
}

func TestHealthCheck(t *testing.T) {
	type args struct {
		hc health.HealthClient
		d  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *grpshuffleClient.HealthCheckResponse
		wantErr bool
	}{
		{
			name: "Return SERVING",
			args: args{
				hc: &mockHealthClient{
					mockCheck: func(ctx context.Context, in *health.HealthCheckRequest) (health.HealthCheckResponse, error) {
						return health.HealthCheckResponse{
							Status: 1,
						}, nil
					},
				},
				d: 3 * time.Second,
			},
			want: &grpshuffleClient.HealthCheckResponse{
				Status: "SERVING",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := grpshuffleClient.HealthCheck(&tt.args.hc)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("actual %s, want %s", actual, tt.want)
			}

		})
	}
}
