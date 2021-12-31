package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

const portNumber = 13333

func extractFields(fullMethod string, req interface{}) map[string]interface{} {
	ret := make(map[string]interface{})

	switch args := req.(type) {
	case *grpshuffle.ShuffleRequest:
		ret["Partition"] = args.Partition
		ret["Targets"] = args.Targets
	default:
		return nil
	}

	return ret
}

func main() {
	kep := keepalive.EnforcementPolicy{
		MinTime: 60 * time.Second,
	}

	log, _ := zap.NewProduction()
	serv := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kep),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)),
				grpc_zap.StreamServerInterceptor(log),
				grpc_prometheus.StreamServerInterceptor,
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)),
				grpc_zap.UnaryServerInterceptor(log),
				grpc_prometheus.UnaryServerInterceptor,
			),
		),
	)

	grpshuffle.RegisterComputeServer(serv, &Server{})
	health.RegisterHealthServer(serv, &healthServer{})
	grpc_prometheus.Register(serv)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":8081", nil); err != nil {
			panic(err)
		}
	}()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}

	err = serv.Serve(l)
	if err != nil {
		fmt.Println("failed to serve:", err)
		os.Exit(1)
	}
}
