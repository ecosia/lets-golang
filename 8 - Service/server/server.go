package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dohe/lets-golang/colorpoint"
)

const (
	addr = "0.0.0.0:9876"
)

func newHandler(client *http.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := colorpoint.Get(client)
		jsonRes, err := json.Marshal(res) // PYTHON: json.dumps()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	}
}

func Serve() {
	client := http.DefaultClient
	http.HandleFunc("/", newHandler(client)) // PYTHON: No built-in routing
	log.Printf("Starting server: http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil)) // PYTHON: http.server's serve_forever()
}
