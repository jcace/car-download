package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "CAR Downloader",
		Usage: "a simple utility to download carfiles from a content data source",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
