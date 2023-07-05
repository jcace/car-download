package main

type Downloader interface {
	DownloadFile(url string, authToken *string, outputLocation string) (string, error)
}

type AriaDownloader struct {
}

func (a AriaDownloader) DownloadFile(url string, authToken *string, outputLocation string) (string, error) {
	return "", nil
}
