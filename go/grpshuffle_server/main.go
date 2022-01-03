package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	grpshuffleServer "github.com/korosuke613/grpshuffle/go/grpshuffle_server/lib"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

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
	app := &cli.App{
		Name:  "grpshuffle-server",
		Usage: "Server of groshuffle",
		Action: func(c *cli.Context) error {
			kep := keepalive.EnforcementPolicy{
				MinTime: 60 * time.Second,
			}

			logger, _ := zap.NewProduction()
			serv := grpc.NewServer(
				grpc.KeepaliveEnforcementPolicy(kep),
				grpc.StreamInterceptor(
					grpc_middleware.ChainStreamServer(
						grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)),
						grpc_zap.StreamServerInterceptor(logger),
						grpc_prometheus.StreamServerInterceptor,
					),
				),
				grpc.UnaryInterceptor(
					grpc_middleware.ChainUnaryServer(
						grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)),
						grpc_zap.UnaryServerInterceptor(logger),
						grpc_prometheus.UnaryServerInterceptor,
					),
				),
			)

			grpshuffle.RegisterComputeServer(serv, &grpshuffleServer.Server{})
			health.RegisterHealthServer(serv, &grpshuffleServer.HealthServer{})
			grpc_prometheus.Register(serv)

			fmt.Println(grpshuffleServer.PrometheusEnable)
			if grpshuffleServer.PrometheusEnable {
				http.Handle("/metrics", promhttp.Handler())
				prometheusAddr := fmt.Sprintf(":%v", grpshuffleServer.PrometheusPort)
				go func() {
					log.Printf("Launch Prometheus server %v...", grpshuffleServer.PrometheusPort)
					if err := http.ListenAndServe(prometheusAddr, nil); err != nil {
						panic(err)
					}
				}()
			}

			log.Printf("Launch grpshuffle server %v...", grpshuffleServer.Port)
			l, err := net.Listen("tcp", fmt.Sprintf(":%d", grpshuffleServer.Port))
			if err != nil {
				return fmt.Errorf("failed to listen:%v", err)
			}

			err = serv.Serve(l)
			if err != nil {
				return fmt.Errorf("failed to serve:%v", err)
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"P"},
				Usage:       "`PORT` of server",
				EnvVars:     []string{"GRPSHUFFLE_PORT"},
				Value:       13333,
				Destination: &grpshuffleServer.Port,
			},
			&cli.BoolFlag{
				Name:        "prometheus-enable",
				Usage:       "With this option, serve Prometheus",
				EnvVars:     []string{"GRPSHUFFLE_PROMETHEUS_ENABLE"},
				Value:       false,
				Destination: &grpshuffleServer.PrometheusEnable,
			},
			&cli.IntFlag{
				Name:        "prometheus-port",
				Usage:       "`PORT` of prometheus",
				EnvVars:     []string{"GRPSHUFFLE_PROMETHEUS_PORT"},
				Value:       8081,
				Destination: &grpshuffleServer.PrometheusPort,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
