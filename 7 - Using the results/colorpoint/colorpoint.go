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
	Color string
	Point struct {
		X int
		Y int
	}
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

func addData(result interface{}, cpResult *Result) {
	if result == nil {
		return
	}
	switch res := result.(type) {
	case *api.HexbotResponse:
		if len(res.Colors) > 0 {
			cpResult.Color = res.Colors[0].Value
		}
	case *api.VexbotResponse:
		if len(res.Vectors) > 0 {
			cpResult.Point = res.Vectors[0].A
		}
	default:
		fmt.Printf("Unrecognized response type: %+v\n", res)
	}
}

func Get() Result {
	results := make(chan jobs.JobResult)
	var startedJobs []jobs.Job

	client := http.DefaultClient
	h := jobs.NewHexBotJob(client)
	startedJobs = append(startedJobs, h)
	go h.Start(results)

	v := jobs.NewVexBotJob(client)
	startedJobs = append(startedJobs, v)
	go v.Start(results)

	var colorPointResult Result

	result, canceled := waitForResponse(startedJobs, results)
	if canceled {
		return colorPointResult
	}
	addData(result, &colorPointResult)

	result, canceled = waitForResponse(startedJobs, results)
	if canceled {
		return colorPointResult
	}
	addData(result, &colorPointResult)
	return colorPointResult
}
