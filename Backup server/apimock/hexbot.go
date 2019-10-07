package apimock

import (
	"encoding/json"
	"net/http"
)

type HexbotValue struct {
	Value string
}

type Hexbot struct {
	Colors []HexbotValue
}

func HexbotHandler(w http.ResponseWriter, r *http.Request) {
	value := HexbotValue{Value: "#D288E9"}
	colors := Hexbot{Colors: []HexbotValue{value}}
	js, err := json.Marshal(colors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
