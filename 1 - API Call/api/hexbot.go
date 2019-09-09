package api // PYTHON: This is like python modules

import (
	"encoding/json"
	"net/http"
)

const (
	HexbotURL = "https://api.noopschallenge.com/hexbot"
)

type HexbotResponse struct { // PYTHON: You'd probably do that with a (data)class
	Colors []struct {
		Value string
	}
}

// CallHexbot calls the hexbot API and parses the returned JSON into a HexbotResponse struct
func CallHexbot() (*HexbotResponse, error) {
	resp, err := http.Get(HexbotURL) // PYTHON: Similar to urllib / requests
	if err != nil {                  // PYTHON: Error handling is done with an error object not a raise/except logic
		return nil, err // PYTHON: None
	}
	return unmarshalHexbotResponse(resp)
}

func unmarshalHexbotResponse(resp *http.Response) (*HexbotResponse, error) {
	var parsed HexbotResponse
	err := json.NewDecoder(resp.Body).Decode(&parsed) // PYTHON: This would just be json.loads(body)
	return &parsed, err                               // PYTHON: Has no explicit pointers
}
