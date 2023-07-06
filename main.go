package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "CAR Downloader",
		Usage: "a simple utility to download carfiles from a content data source",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Value:   "./",
				Usage:   "output directory for downloaded files",
			},
			&cli.StringFlag{
				Name:    "aria-uri",
				Aliases: []string{"a"},
				Value:   "ws://localhost:6800/jsonrpc",
				Usage:   "aria2c RPC URI",
			},
			&cli.StringFlag{
				Name:     "ddm-api",
				Aliases:  []string{"d"},
				Required: true,
				Usage:    "DDM URI to fetch contents from",
			},
			&cli.StringFlag{
				Name:     "ddm-token",
				Aliases:  []string{"t"},
				Required: true,
				Usage:    "DDM self-service token",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("starting downloader")

			cfg, err := CreateConfig(c)
			if err != nil {
				log.Fatal(err)
			}

			d := NewAriaDownloader(cfg.AriaUri)
			ddm := NewDDMContentSource(cfg.DDMUri, cfg.DDMToken)

			RunDaemon(d, ddm, cfg)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
