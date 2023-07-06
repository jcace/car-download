package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Content struct {
	PieceCID        string "json:piece_cid"
	ContentLocation string "json:content_location"
}

// A content source is a location that returns a list of available files that can be downloaded and later used for deals
type ContentSource interface {
	GetContentList() ([]Content, error)
}

type DDMContentSource struct {
	baseUrl string
	authKey string
}

func NewDDMContentSource(baseUrl string, authKey string) *DDMContentSource {
	return &DDMContentSource{baseUrl: baseUrl, authKey: authKey}
}

type ResponseDDMAvailableContent struct {
	PayloadCID      string `json:"payload_cid"`
	PieceCID        string `json:"piece_cid"`
	Size            uint64 `json:"size"`
	PaddedSize      uint64 `json:"padded_size"`
	ContentLocation string `json:"content_location"`
}

func (d *DDMContentSource) GetContentList() ([]Content, error) {
	req, err := http.NewRequest(http.MethodGet, d.baseUrl+"/api/v1/self-service/available-contents", nil)
	if err != nil {
		return nil, fmt.Errorf("could not construct http request %v", err)
	}

	req.Header.Set("X-DELTA-AUTH", d.authKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute http request %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error in http call %d : %s", resp.StatusCode, body)
	}

	if err != nil {
		return nil, fmt.Errorf("could not read response body %v", err)
	}

	var availableContents []ResponseDDMAvailableContent
	err = json.Unmarshal(body, &availableContents)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response %v", err)
	}

	var contents []Content

	for _, availableContent := range availableContents {
		if availableContent.ContentLocation == "" {
			continue
		}
		contents = append(contents, Content{
			PieceCID:        availableContent.PieceCID,
			ContentLocation: availableContent.ContentLocation,
		})
	}

	return contents, nil
}
