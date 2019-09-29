package jobs

import (
	"context"
	"net/http"

	"github.com/dohe/lets-golang/api"
)

const (
	vexBotID = "vexbot"
)

// VexBotJob is the base type for this job
type VexBotJob struct {
	cancelContext context.Context
	cancelFunc    context.CancelFunc
	httpClient    *http.Client
}

// ID returns the job's ID
func (v VexBotJob) ID() string {
	return vexBotID
}

// Cancel cancels any outstanding requests
func (v VexBotJob) Cancel() {
	v.cancelFunc()
}

// Start runs the job
func (v VexBotJob) Start(done chan JobResult) {
	resp, err := api.CallVexbot(v.cancelContext, v.httpClient)
	done <- JobResult{ID: vexBotID, Result: resp, Err: err}
}

// NewVexBotJob creates a new job that returns a vexbot response
func NewVexBotJob(httpClient *http.Client) VexBotJob {
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	return VexBotJob{
		cancelContext,
		cancelFunc,
		httpClient,
	}
}
