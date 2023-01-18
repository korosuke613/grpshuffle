package grpshuffle_server

import (
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

func Shuffle(req *grpshuffle.ShuffleRequest) (*grpshuffle.ShuffleResponse, error) {
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
	average := int(uint64(len(req.Targets)) / req.Divide)
	remainder := int(uint64(len(req.Targets)) % req.Divide)

	var sliceSize int
	loopCount := 0
	for i := 0; i < len(req.Targets); i += sliceSize {
		if loopCount < remainder {
			sliceSize = average + 1
		} else {
			sliceSize = average
		}
		endCursor := i + sliceSize

		if endCursor > len(req.Targets) {
			endCursor = len(req.Targets)
		}
		tmp := shuffledTargets[i:endCursor]
		slicedTargets = append(slicedTargets, &grpshuffle.Combination{Targets: tmp})
		loopCount += 1
	}

	return &grpshuffle.ShuffleResponse{
		Combinations: slicedTargets,
	}, nil
}
