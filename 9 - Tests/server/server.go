package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dohe/lets-golang/colorpoint"
)

const (
	port = "0.0.0.0:9876"
)

func newHandler(client *http.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := colorpoint.Get(client)
		jsonRes, err := json.Marshal(res)
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
	http.HandleFunc("/", newHandler(client))
	log.Printf("Starting server: http://%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
