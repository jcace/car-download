package main

import (
	"fmt"
)

func RunDaemon(d Downloader, c ContentSource, cfg Config) {
	for {
		// Get the list of content available
		contentList, err := c.GetContentList()
		if err != nil {
			fmt.Println(err)
		}

		// For each content, download it
		for _, content := range contentList {
			fmt.Printf("downloading %s to %s\n", content.PieceCID, content.ContentLocation)
			err := d.DownloadFile(content.PieceCID, nil, cfg.OutDir, content.PieceCID+".car")
			if err != nil {
				fmt.Printf("error downloading %s: %s\n", content.PieceCID, err)
			}
		}
	}
}
