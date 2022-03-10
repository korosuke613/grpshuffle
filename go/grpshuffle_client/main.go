package main

import (
	"encoding/json"
	"fmt"
	"github.com/korosuke613/grpshuffle/go/grpshuffle"
	grpshuffleClient "github.com/korosuke613/grpshuffle/go/grpshuffle_client/lib"

	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var divide int
	app := &cli.App{
		Name:  "grpshuffle-client",
		Usage: "Client of groshuffle",
		Commands: []*cli.Command{
			{
				Name:  "shuffle",
				Usage: "shuffle",
				Flags: append([]cli.Flag{
					&cli.IntFlag{
						Name:        "divide",
						Aliases:     []string{"d"},
						Usage:       "Number to divide",
						EnvVars:     []string{"GRPSHUFFLE_SHUFFLE_DIVIDE"},
						Required:    true,
						Destination: &divide,
					},
				}, newGlobalFlags()...),
				ArgsUsage: "DIVIDE TARGET1 TARGET2 ...",
				Action: func(c *cli.Context) error {
					conn, err := grpshuffleClient.Connect(grpshuffleClient.Host, grpshuffleClient.Port, grpshuffleClient.NoTLS)
					if err != nil {
						return err
					}
					defer grpshuffleClient.CloseConnect(conn)

					cc := grpshuffle.NewComputeClient(conn)
					rawResult, err := grpshuffleClient.Shuffle(&cc, uint64(divide), c.Args().Slice())
					if err != nil {
						return err
					}

					result, err := json.MarshalIndent(rawResult, "", "  ")
					if err != nil {
						return err
					}

					fmt.Println(string(result))
					return nil
				},
			},
			{
				Name:  "health",
				Usage: "health check server",
				Action: func(c *cli.Context) error {
					conn, err := grpshuffleClient.Connect(grpshuffleClient.Host, grpshuffleClient.Port, grpshuffleClient.NoTLS)
					if err != nil {
						return err
					}
					defer grpshuffleClient.CloseConnect(conn)

					hc := grpc_health_v1.NewHealthClient(conn)
					rawResult, err := grpshuffleClient.HealthCheck(&hc)
					if err != nil {
						return err
					}
					result, err := json.MarshalIndent(rawResult, "", "  ")
					if err != nil {
						return err
					}

					fmt.Println(string(result))
					return nil
				},
				Flags: append([]cli.Flag{}, newGlobalFlags()...),
			},
			{
				Name: "http-serve",
				Action: func(c *cli.Context) error {
					grpshuffleClient.HttpServe(8080)
					return nil
				},
				Flags: append([]cli.Flag{}, newGlobalFlags()...),
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newGlobalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"H"},
			Usage:       "Host address of server",
			EnvVars:     []string{"GRPSHUFFLE_HOST"},
			Value:       "localhost",
			Destination: &grpshuffleClient.Host,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"P"},
			Usage:       "Port of server",
			EnvVars:     []string{"GRPSHUFFLE_PORT"},
			Value:       13333,
			Destination: &grpshuffleClient.Port,
		},
		&cli.BoolFlag{
			Name:        "no-tls",
			Usage:       "If this flag is enabled, TLS is not used.",
			EnvVars:     []string{"GRPSHUFFLE_NO_TLS"},
			Value:       false,
			Destination: &grpshuffleClient.NoTLS,
		},
	}
}
