package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dohe/lets-golang/api"
	"github.com/dohe/lets-golang/jobs"
)

const (
	timeout = 2 * time.Second
)

func waitForResponse(startedJobs []jobs.Job, results chan jobs.JobResult) bool {
	select {
	case result := <-results:
		if result.Err != nil {
			fmt.Printf("%s failed: %s\n", result.ID, result.Err.Error())
		} else {
			switch res := result.Result.(type) {
			case *api.HexbotResponse:
				fmt.Printf("Hexbot successful: %+v\n", res)
			case *api.VexbotResponse:
				fmt.Printf("Vexbot successful: %+v\n", res)
			default:
				fmt.Printf("Unknown successful: %+v\n", res)
			}
		}
		return false
	case <-time.After(timeout): // PYTHON: something similar can be achieved with Queue.get(timeout=timeout)
		for _, job := range startedJobs {
			job.Cancel()
			fmt.Printf("Canceled %s\n", job.ID())
		}
		return true
	}
}

func main() {
	fmt.Println("Start")
	results := make(chan jobs.JobResult)
	var startedJobs []jobs.Job

	client := http.DefaultClient
	h := jobs.NewHexBotJob(client)
	startedJobs = append(startedJobs, h)
	go h.Start(results)

	v := jobs.NewVexBotJob(client)
	startedJobs = append(startedJobs, v)
	go v.Start(results)

	canceled1 := waitForResponse(startedJobs, results)
	canceled2 := waitForResponse(startedJobs, results)
	if canceled1 || canceled2 {
		waitForResponse(startedJobs, results)
	}

	fmt.Println("Exit")
}
