package main

import (
	"context"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"reflect"
	"testing"
)

func TestServer_Shuffle(t *testing.T) {
	type fields struct {
		UnimplementedComputeServer grpshuffle.UnimplementedComputeServer
	}
	type args struct {
		ctx context.Context
		req *grpshuffle.ShuffleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpshuffle.ShuffleResponse
		wantErr bool
	}{
		{
			name:   "There are no omissions in the items that come back.",
			fields: fields{UnimplementedComputeServer: grpshuffle.UnimplementedComputeServer{}},
			args: args{
				ctx: context.Background(),
				req: &grpshuffle.ShuffleRequest{
					Targets:    []string{"a", "b", "c", "d", "e"},
					Partition:  1,
					Sequential: true,
				},
			},
			want: &grpshuffle.ShuffleResponse{
				Combinations: []*grpshuffle.Combination{
					{Targets: []string{"a", "b", "c", "d", "e"}},
				},
			},
		},
		{
			name:   "Split into two.",
			fields: fields{UnimplementedComputeServer: grpshuffle.UnimplementedComputeServer{}},
			args: args{
				ctx: context.Background(),
				req: &grpshuffle.ShuffleRequest{
					Targets:    []string{"a", "b", "c", "d", "e"},
					Partition:  2,
					Sequential: true,
				},
			},
			want: &grpshuffle.ShuffleResponse{
				Combinations: []*grpshuffle.Combination{
					{Targets: []string{"a", "b", "c"}},
					{Targets: []string{"d", "e"}},
				},
			},
		},
		{
			name:   "Split into three.",
			fields: fields{UnimplementedComputeServer: grpshuffle.UnimplementedComputeServer{}},
			args: args{
				ctx: context.Background(),
				req: &grpshuffle.ShuffleRequest{
					Targets:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
					Partition:  3,
					Sequential: true,
				},
			},
			want: &grpshuffle.ShuffleResponse{
				Combinations: []*grpshuffle.Combination{
					{Targets: []string{"a", "b", "c"}},
					{Targets: []string{"d", "e", "f"}},
					{Targets: []string{"g", "h", "i"}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedComputeServer: tt.fields.UnimplementedComputeServer,
			}
			got, err := s.Shuffle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() got = %v, want %v", got, tt.want)
			}
		})
	}
}
