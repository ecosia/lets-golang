package main

import (
	"log" // PYTHON: Similar to python's logging
	"os"

	"github.com/dohe/lets-golang/api"
	"github.com/dohe/lets-golang/jobs"
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func waitForResponse(results chan jobs.JobResult) {
	select {
	case result := <-results: // PYTHON: You could build something similar with queue and threads
		if result.Err != nil {
			logger.Printf("%s failed: %s\n", result.ID, result.Err.Error())
		} else {
			switch res := result.Result.(type) {
			case *api.HexbotResponse:
				logger.Printf("Hexbot successful: %+v\n", res)
			case *api.VexbotResponse:
				logger.Printf("Vexbot successful: %+v\n", res)
			default:
				logger.Printf("Unknown successful: %+v\n", res)
			}
		}
	}
}

func main() {
	logger.Println("Start")
	results := make(chan jobs.JobResult)
	numCalls := 0

	h := jobs.NewHexBotJob()
	numCalls++
	go h.Start(results)

	v := jobs.NewVexBotJob()
	numCalls++
	go v.Start(results)

	for numCalls > 0 { // PYTHON: this would be a while loop
		waitForResponse(results)
		numCalls--
	}
	logger.Println("Exit")
}
