package api

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	HexbotURL = "https://api.noopschallenge.com/hexbot"
	// HexbotURL = "http://slowwly.robertomurray.co.uk/delay/1000/url/https://api.noopschallenge.com/hexbot"
)

type HexbotResponse struct {
	Colors []struct {
		Value string
	}
}

// CallHexbot calls the hexbot API and parses the returned JSON into a HexbotResponse struct
func CallHexbot(cancelContext context.Context, client *http.Client) (*HexbotResponse, error) {
	req, err := http.NewRequestWithContext(cancelContext, "GET", HexbotURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return unmarshalHexbotResponse(resp)
}

func unmarshalHexbotResponse(resp *http.Response) (*HexbotResponse, error) {
	var parsed HexbotResponse
	err := json.NewDecoder(resp.Body).Decode(&parsed)
	return &parsed, err
}
