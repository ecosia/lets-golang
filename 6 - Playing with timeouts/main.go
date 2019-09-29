package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dohe/lets-golang/api"
	"github.com/dohe/lets-golang/jobs"
)

const (
	timeout = 2 * time.Second
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds) // PYTHON: Similar to python's logging
)

func waitForResponse(jobs []jobs.Job, results chan jobs.JobResult) bool {
	select {
	case result := <-results:
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
		return false
	case <-time.After(timeout):
		for _, job := range jobs {
			job.Cancel()
			logger.Printf("Canceled %s\n", job.ID())
		}
		return true
	}
}

func main() {
	logger.Println("Start")
	results := make(chan jobs.JobResult)
	var startedJobs []jobs.Job

	client := http.DefaultClient
	h := jobs.NewHexBotJob(client)
	startedJobs = append(startedJobs, h)
	go h.Start(results)

	v := jobs.NewVexBotJob(client)
	startedJobs = append(startedJobs, v)
	go v.Start(results)

	numReceived := 0
	for numReceived < len(startedJobs) {
		canceled := waitForResponse(startedJobs, results)
		if !canceled {
			numReceived++
		}
	}
	logger.Println("Exit")
}
