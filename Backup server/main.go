package main

import (
	"log"
	"net/http"

	"github.com/dohe/lets-golang/apimock"
)

const (
	addr = "0.0.0.0:3000"
)

func main() {
	http.HandleFunc("/hexbot", apimock.HexbotHandler)
	http.HandleFunc("/vexbot", apimock.VexbotHandler)
	log.Printf("Starting server: http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
