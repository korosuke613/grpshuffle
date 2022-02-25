package grpshuffle_server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"math/rand"
	"time"

	"github.com/korosuke613/grpshuffle/go/grpshuffle"
)

// Server is implemented ComputeServer
type Server struct {
	// 将来 proto ファイルに RPC が追加されてインタフェースが拡張された際、
	// ビルドエラーになるのを防止する仕組み。
	grpshuffle.UnimplementedComputeServer
}

// インタフェースが実装できていることをコンパイル時に確認するおまじない
var _ grpshuffle.ComputeServer = &Server{}

// Shuffle is gRPC server
func (s *Server) Shuffle(ctx context.Context, req *grpshuffle.ShuffleRequest) (*grpshuffle.ShuffleResponse, error) {
	shuffledTargets := make([]string, len(req.Targets))
	copy(shuffledTargets, req.Targets)

	if !req.Sequential {
		// randomly swap targets.
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(req.Targets), func(i, j int) {
			shuffledTargets[i], shuffledTargets[j] = shuffledTargets[j], shuffledTargets[i]
		})
	}

	if req.Divide >= uint64(len(req.Targets)) {
		return nil, status.Errorf(codes.InvalidArgument, "Must have `divide` >= `targets` num.")
	}

	// split targets by the number of divide.
	slicedTargets := make([]*grpshuffle.Combination, 0)
	sliceSize := int(math.Ceil(float64(len(req.Targets)) / float64(req.Divide)))
	for i := 0; i < len(req.Targets); i += sliceSize {
		endCursor := i + sliceSize
		if endCursor > len(req.Targets) {
			endCursor = len(req.Targets)
		}
		tmp := shuffledTargets[i:endCursor]
		slicedTargets = append(slicedTargets, &grpshuffle.Combination{Targets: tmp})
	}

	return &grpshuffle.ShuffleResponse{
		Combinations: slicedTargets,
	}, nil
}
