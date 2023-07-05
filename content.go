package main

type Content struct {
	PieceCID        string "json:piece_cid"
	ContentLocation string "json:content_location"
}

// A content source is a location that returns a list of available files that can be downloaded and later used for deals
type ContentSource interface {
	GetContentList(url string, authToken *string) ([]Content, error)
}
