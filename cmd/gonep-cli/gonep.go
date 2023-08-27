package main

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	gonep2 "github.com/maxihafer/gonep/pkg/gonep"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var client *gonep2.Client

	app := &cli.App{
		Name:  "gonep",
		Usage: "read nepviewer data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Category: "Authentication",
				Usage:    "The token used to authenticate to the NepViewer API. If this is set, the username and password flags are ignored.",
				EnvVars:  []string{"NEP_VIEWER_TOKEN"},
			},
			&cli.StringFlag{
				Name:     "user",
				Aliases:  []string{"u"},
				Category: "Authentication",
				Usage:    "The username used to authenticate against the NepViewer API. TOKEN takes precedence over user/password authentication.",
				EnvVars:  []string{"NEP_VIEWER_USER"},
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Category: "Authentication",
				Usage:    "The password used to authenticate against the NepViewer API. Token takes precedence over user/password authentication.",
				EnvVars:  []string{"NEP_VIEWER_PASSWORD"},
			},
		},
		Before: func(ctx *cli.Context) error {
			var opts []gonep2.ClientOption
			var err error

			if token := ctx.String("token"); token == "" {
				username := ctx.String("user")
				password := ctx.String("password")

				if username == "" {
					return errors.New("username not set. Either token oder user/password authentication must be setup")
				}

				if password == "" {
					return errors.New("password not set. Either token oder user/password authentication must be setup")
				}

				opts = append(opts, gonep2.WithUserPassword(username, password))
			} else {
				opts = append(opts, gonep2.WithToken(token))
			}

			client, err = gonep2.NewClient(opts...)
			return err
		},
		Commands: []*cli.Command{
			{
				Name:     "plant",
				Category: "Plants",
				Usage:    "query plant information",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list photovoltaic plants",
						Action: func(ctx *cli.Context) error {
							plants, err := client.Plants().List(ctx.Context)
							if err != nil {
								return err
							}

							t := table.NewWriter()
							t.AppendHeader(table.Row{"Site ID", "Name", "Gateways"})
							for _, plant := range plants {
								for _, gw := range plant.Gateways {
									t.AppendRow(table.Row{plant.Sid, plant.SiteName, gw}, table.RowConfig{AutoMerge: true})
								}
							}
							fmt.Println(t.Render())

							return nil
						},
					},
					{
						Name:   "get",
						Usage:  "get the status of a photovoltaic plant",
						Action: func(ctx *cli.Context) error { return errors.New("not yet implemented") },
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
