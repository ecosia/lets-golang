package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
	"github.com/dohe/lets-golang/jobs"
)

func waitForResponse(results chan jobs.JobResult) {
	result := <-results // PYTHON: You could build something similar with queue and threads
	if result.Err != nil {
		fmt.Printf("%s failed: %s\n", result.ID, result.Err.Error())
	} else {
		switch res := result.Result.(type) { // PYTHON: similar to isinstance
		case *api.HexbotResponse: // PYTHON: if/elif/else
			fmt.Printf("Hexbot successful: %+v\n", res)
		case *api.VexbotResponse:
			fmt.Printf("Vexbot successful: %+v\n", res)
		default:
			fmt.Printf("Unknown successful: %+v\n", res)
		}
	}
}

func main() {
	fmt.Println("Start")
	results := make(chan jobs.JobResult)

	h := jobs.NewHexBotJob()
	go h.Start(results)

	v := jobs.NewVexBotJob()
	go v.Start(results)

	waitForResponse(results)
	waitForResponse(results)

	fmt.Println("Exit")
}
