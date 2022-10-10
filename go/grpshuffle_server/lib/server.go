package grpshuffle_server

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/korosuke613/grpshuffle/gen/korosuke613/grpshuffle/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

// Server is implemented ComputeServer
type Server struct {
	Logger *zap.SugaredLogger
}

// Shuffle is gRPC server
func (s *Server) Shuffle(ctx context.Context, req *connect.Request[grpshufflev1.ShuffleRequest]) (*connect.Response[grpshufflev1.ShuffleResponse], error) {
	if s.Logger != nil {
		//msgJson, err := json.Marshal(req.Msg)
		//if err != nil {
		//	s.Logger.Errorw("json marshal error",
		//		"err", err,
		//	)
		//}
		fmt.Printf("%v", req.Msg)
		s.Logger.Infow("calling",
			"header", req.Header(),
			"spec", req.Spec(),
			"peer", req.Peer(),
			"request", &req.Msg,
		)
	}

	shuffledTargets := make([]string, len(req.Msg.Targets))
	copy(shuffledTargets, req.Msg.Targets)

	if !req.Msg.Sequential {
		// randomly swap targets.
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(req.Msg.Targets), func(i, j int) {
			shuffledTargets[i], shuffledTargets[j] = shuffledTargets[j], shuffledTargets[i]
		})
	}

	if req.Msg.Divide >= uint64(len(req.Msg.Targets)) {
		return nil, status.Errorf(codes.InvalidArgument, "Must have `divide` >= `targets` num.")
	}

	// split targets by the number of divide.
	slicedTargets := make([]*grpshufflev1.Combination, 0)
	average := int(uint64(len(req.Msg.Targets)) / req.Msg.Divide)
	remainder := int(uint64(len(req.Msg.Targets)) % req.Msg.Divide)

	var sliceSize int
	loopCount := 0
	for i := 0; i < len(req.Msg.Targets); i += sliceSize {
		if loopCount < remainder {
			sliceSize = average + 1
		} else {
			sliceSize = average
		}
		endCursor := i + sliceSize

		if endCursor > len(req.Msg.Targets) {
			endCursor = len(req.Msg.Targets)
		}
		tmp := shuffledTargets[i:endCursor]
		slicedTargets = append(slicedTargets, &grpshufflev1.Combination{Targets: tmp})
		loopCount += 1
	}

	res := connect.NewResponse(&grpshufflev1.ShuffleResponse{
		Combinations: slicedTargets,
	})

	return res, nil
}
