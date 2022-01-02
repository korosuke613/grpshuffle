package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	host string
	port int
)

func main() {
	var partition int
	app := &cli.App{
		Name:  "grpshuffle-client",
		Usage: "Client of groshuffle",
		Commands: []*cli.Command{
			{
				Name:  "shuffle",
				Usage: "shuffle",
				Flags: append([]cli.Flag{
					&cli.IntFlag{
						Name:        "partition",
						Aliases:     []string{"p"},
						Usage:       "Number to divide",
						EnvVars:     []string{"GRPSHUFFLE_SHUFFLE_PARTITION"},
						Required:    true,
						Destination: &partition,
					},
				}, newGlobalFlags()...),
				ArgsUsage: "PARTITION TARGET1 TARGET2 ...",
				Action: func(c *cli.Context) error {
					if partition > c.Args().Len() {
						return fmt.Errorf("the number of TARGET must be greater than partition")
					}

					conn, err := connect(host, port)
					if err != nil {
						return err
					}
					defer func(conn *grpc.ClientConn) {
						err := conn.Close()
						if err != nil {
							log.Fatal(err)
						}
					}(conn)

					err = callShuffle(conn, int32(partition), c.Args().Slice())
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "health",
				Usage: "health check server",
				Action: func(c *cli.Context) error {
					conn, err := connect(host, port)
					if err != nil {
						return err
					}
					defer func(conn *grpc.ClientConn) {
						err := conn.Close()
						if err != nil {
							log.Fatal(err)
						}
					}(conn)
					err = callHealth(conn)
					if err != nil {
						return err
					}
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
			Destination: &host,
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"P"},
			Usage:       "Port of server",
			EnvVars:     []string{"GRPSHUFFLE_PORT"},
			Value:       13333,
			Destination: &port,
		},
	}
}
