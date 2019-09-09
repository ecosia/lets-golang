package api

import (
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
func CallVexbot() (*VexbotResponse, error) {
	resp, err := http.Get(VexbotURL)
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
