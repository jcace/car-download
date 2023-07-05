package main

import "github.com/urfave/cli/v2"

type Config struct {
	OutDir   string
	AriaUri  string
	DDMUri   string
	DDMToken string
}

func CreateConfig(cctx *cli.Context) (Config, error) {
	config := Config{
		OutDir:   cctx.String("out"),
		AriaUri:  cctx.String("aria-uri"),
		DDMUri:   cctx.String("ddm-api"),
		DDMToken: cctx.String("ddm-token"),
	}

	return config, nil
}
