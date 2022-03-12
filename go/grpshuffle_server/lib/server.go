package grpshuffle_server

import (
	"context"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"time"
)

// Server is implemented ComputeServer
type Server struct {
	// 将来 proto ファイルに RPC が追加されてインタフェースが拡張された際、
	// ビルドエラーになるのを防止する仕組み。
	grpshuffle.UnimplementedComputeServer
}

// インタフェースが実装できていることをコンパイル時に確認するおまじない
var _ grpshuffle.ComputeServer = &Server{}

// Shuffle is shuffle
func (s *Server) Shuffle(ctx context.Context, req *grpshuffle.ShuffleRequest) (*grpshuffle.ShuffleResponse, error) {
	return Shuffle(req)
}

// RepeatShuffle is repeat shuffle.
func (s *Server) RepeatShuffle(req *grpshuffle.RepeatShuffleRequest, stream grpshuffle.Compute_RepeatShuffleServer) error {
	times := uint64(0)

	for {
		select {
		// Exit when the client cancels the request.
		case <-stream.Context().Done():
			return nil
		// Wait for `req.Interval` second
		case <-time.After(time.Duration(req.Interval) * time.Second):
		}

		// Exit after the number of attempts is exceeded.
		if req.Times != 0 && req.Times < times {
			return nil
		}

		result, err := Shuffle(req.ShuffleRequest)
		if err != nil {
			return err
		}
		if err := stream.Send(result); err != nil {
			return err
		}

		times += 1
	}
}
