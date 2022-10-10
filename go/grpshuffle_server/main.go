package main

import (
	"fmt"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/korosuke613/grpshuffle/gen/korosuke613/grpshuffle/v1/grpshufflev1connect"
	grpshuffleServer "github.com/korosuke613/grpshuffle/go/grpshuffle_server/lib"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "grpshuffle-server",
		Usage: "Server of groshuffle",
		Action: func(c *cli.Context) error {
			mux := http.NewServeMux()
			logger, _ := zap.NewProduction()
			suger := logger.Sugar()

			path, handler := grpshufflev1connect.NewComputeServiceHandler(
				&grpshuffleServer.Server{Logger: suger},
			)
			mux.Handle(path, handler)

			checker := grpchealth.NewStaticChecker(
				grpshufflev1connect.ComputeServiceName,
			)
			mux.Handle(grpchealth.NewHandler(checker))

			reflector := grpcreflect.NewStaticReflector(
				grpshufflev1connect.ComputeServiceName,
				"grpc.health.v1.Health",
			)
			mux.Handle(grpcreflect.NewHandlerV1(reflector))
			mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

			addr := fmt.Sprintf(":%d", grpshuffleServer.Port)
			log.Printf("Launch grpshuffle server %v...", addr)
			err := http.ListenAndServe(
				addr,
				h2c.NewHandler(mux, &http2.Server{}),
			)
			if err != nil {
				return err
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
