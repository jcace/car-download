package main

import (
	"fmt"
	"path"
	"time"
)

const CONTENT_GET_INTERVAL = 10 * time.Minute

func RunDaemon(d Downloader, c ContentSource, cfg Config) {
	for {
		// Get the list of content available
		contentList, err := c.GetContentList()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("got a list of %d contents to download", len(contentList))

		// For each content, download it
		for _, content := range contentList {
			if FileExists(path.Join(cfg.OutDir, content.PieceCID+".car")) {
				fmt.Printf("skipping %s, already downloaded\n", content.PieceCID)
				continue
			}
			fmt.Printf("downloading %s to %s\n", content.ContentLocation, content.PieceCID)
			err = d.DownloadFile(content.ContentLocation, nil, cfg.OutDir, content.PieceCID+".car")
			if err != nil {
				fmt.Printf("error downloading %s: %s\n", content.PieceCID, err)
			}

		}
		fmt.Printf("completed content download. waiting %d minutes before check for new content", CONTENT_GET_INTERVAL/time.Minute)

		// Wait between runs
		time.Sleep(CONTENT_GET_INTERVAL)
	}
}
