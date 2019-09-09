package colorpoint

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type ClientMock struct {
	hexBody string
	vexBody string
	err     error
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	body := ""
	if strings.Contains(req.URL.Path, "hex") {
		body = c.hexBody
	}
	if strings.Contains(req.URL.Path, "vex") {
		body = c.vexBody
	}
	return &http.Response{
		StatusCode: 200,
		Status:     "OK",
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}, c.err
}

func clientMock(hexBody, vexBody string, err error) ClientMock {
	return ClientMock{hexBody, vexBody, err}
}

func TestGet(t *testing.T) {
	hexResponse := `{
		"colors": [
			{ "value": "#FFAA00" }
		]
	}`
	vexResponse := `{
		"vectors":
			[
				{ 
					"a": { "x":1, "y":1 },
					"b": { "x":2, "y":2 }
				}
			]
	}`
	tests := []struct {
		name     string
		hexBody  string
		vexBody  string
		err      error
		expected Result
	}{
		{"Empty response", "", "", nil, Result{}},
		{"Empty json", "{}", "{}", nil, Result{}},
		{"Hex only", hexResponse, "{}", nil, Result{Color: "#FFAA00"}},
		{"Vex only", "{}", vexResponse, nil, Result{Point: Point{X: 1, Y: 1}}},
		{"Default", hexResponse, vexResponse, nil, Result{Color: "#FFAA00", Point: Point{X: 1, Y: 1}}},
		{"Error when calling", hexResponse, vexResponse, errors.New("Ouch"), Result{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := clientMock(test.hexBody, test.vexBody, test.err)
			actual := Get(&client)
			if actual != test.expected {
				t.Errorf("Wanted %+v, got %+v", test.expected, actual)
			}
		})
	}
}
