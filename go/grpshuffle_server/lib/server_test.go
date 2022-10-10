package grpshuffle_server_test

import (
	"context"
	"github.com/bufbuild/connect-go"
	grpshuffleServer "github.com/korosuke613/grpshuffle/go/grpshuffle_server/lib"
	"reflect"
	"testing"

	grpshuffle "github.com/korosuke613/grpshuffle/gen/korosuke613/grpshuffle/v1"
)

func TestServer_Shuffle(t *testing.T) {
	type fields struct{}
	type args struct {
		ctx context.Context
		req *connect.Request[grpshuffle.ShuffleRequest]
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
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[grpshuffle.ShuffleRequest]{
					Msg: &grpshuffle.ShuffleRequest{
						Targets:    []string{"a", "b", "c", "d", "e"},
						Divide:     1,
						Sequential: true,
					},
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
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[grpshuffle.ShuffleRequest]{
					Msg: &grpshuffle.ShuffleRequest{
						Targets:    []string{"a", "b", "c", "d", "e"},
						Divide:     2,
						Sequential: true,
					},
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
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[grpshuffle.ShuffleRequest]{
					Msg: &grpshuffle.ShuffleRequest{
						Targets:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
						Divide:     3,
						Sequential: true,
					},
				},
			},
			want: &grpshuffle.ShuffleResponse{
				Combinations: []*grpshuffle.Combination{
					{Targets: []string{"a", "b", "c", "d"}},
					{Targets: []string{"e", "f", "g"}},
					{Targets: []string{"h", "i", "j"}},
				},
			},
		},
		{
			name:   "Split into four.",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &connect.Request[grpshuffle.ShuffleRequest]{
					Msg: &grpshuffle.ShuffleRequest{
						Targets:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
						Divide:     4,
						Sequential: true,
					},
				},
			},
			want: &grpshuffle.ShuffleResponse{
				Combinations: []*grpshuffle.Combination{
					{Targets: []string{"a", "b", "c"}},
					{Targets: []string{"d", "e", "f"}},
					{Targets: []string{"g", "h"}},
					{Targets: []string{"i", "j"}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &grpshuffleServer.Server{}
			got, err := s.Shuffle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Msg, tt.want) {
				t.Errorf("Shuffle() got = %v, want %v", got.Msg, tt.want)
			}
		})
	}
}
