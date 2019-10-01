package api

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	VexbotURL = "https://api.noopschallenge.com/vexbot"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	A     Point
	B     Point
	Speed int
}

type VexbotResponse struct {
	Vectors []Vector
}

// CallVexbot calls the vexbot API and parses the returned JSON into a VexbotResponse struct
func CallVexbot(cancelContext context.Context, client *http.Client) (*VexbotResponse, error) {
	req, err := http.NewRequestWithContext(cancelContext, "GET", VexbotURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return unmarshalVexbotResponse(resp)
}

func unmarshalVexbotResponse(resp *http.Response) (*VexbotResponse, error) {
	var parsed VexbotResponse
	err := json.NewDecoder(resp.Body).Decode(&parsed)
	return &parsed, err
}
