package colorpoint

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

type Result struct {
	Color string `json:"color,omitempty"`
	Point struct {
		X int
		Y int
	} `json:"point,omitempty"`
}

func waitForResponse(jobs []jobs.Job, results chan jobs.JobResult) (interface{}, bool) {
	select {
	case result := <-results:
		if result.Err != nil {
			fmt.Printf("%s failed: %s\n", result.ID, result.Err.Error())
			return nil, false
		}
		return result.Result, false
	case <-time.After(timeout):
		for _, job := range jobs {
			job.Cancel()
			fmt.Printf("Canceled %s\n", job.ID())
		}
		return nil, true
	}
}

func Get(client *http.Client) Result {
	results := make(chan jobs.JobResult)
	var startedJobs []jobs.Job

	h := jobs.NewHexBotJob(client)
	startedJobs = append(startedJobs, h)
	go h.Start(results)

	v := jobs.NewVexBotJob(client)
	startedJobs = append(startedJobs, v)
	go v.Start(results)

	var colorPointResult Result
	numReceived := 0
	for numReceived < len(startedJobs) {
		result, canceled := waitForResponse(startedJobs, results)
		if canceled {
			continue
		}
		numReceived++
		if result == nil {
			continue
		}
		switch res := result.(type) {
		case *api.HexbotResponse:
			if len(res.Colors) > 0 {
				colorPointResult.Color = res.Colors[0].Value
			}
		case *api.VexbotResponse:
			if len(res.Vectors) > 0 {
				colorPointResult.Point = res.Vectors[0].A
			}
		default:
			fmt.Printf("Unrecognized response type: %+v\n", res)
		}
	}

	return colorPointResult
}
