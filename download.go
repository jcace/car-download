package main

import (
	"fmt"
	"log"

	"github.com/siku2/arigo"
)

type Downloader interface {
	DownloadFile(url string, authToken *string, outDir string, outFilename string) error
}

type AriaDownloader struct {
	c *arigo.Client
}

func NewAriaDownloader(uri string) *AriaDownloader {
	c, err := arigo.Dial(uri, "")

	if err != nil {
		log.Fatalf("error connecting to aria2c: %s", err)
	}
	return &AriaDownloader{
		c: &c,
	}
}

func (a *AriaDownloader) DownloadFile(url string, authToken *string, outDir string, outFilename string) error {
	// TODO: parameterize maxConnections
	status, err := a.c.Download(arigo.URIs(url), &arigo.Options{Dir: outDir, Out: outFilename, MaxConnectionPerServer: 8, AutoFileRenaming: false, AllowOverwrite: false, ConditionalGet: true})
	if err != nil {
		return err
	}

	// download complete
	fmt.Printf("download complete! %d bytes/s\n", status.DownloadSpeed)
	return nil
}
