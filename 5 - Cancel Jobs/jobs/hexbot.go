package jobs

import (
	"context"
	"net/http"

	"github.com/dohe/lets-golang/api"
)

const (
	hexBotID = "hexbot"
)

// HexBotJob is the base type for this job
type HexBotJob struct {
	cancelContext context.Context
	cancelFunc    context.CancelFunc
	httpClient    *http.Client
}

// ID returns the job's ID
func (h HexBotJob) ID() string {
	return hexBotID
}

// Cancel cancels any outstanding requests
func (h HexBotJob) Cancel() {
	h.cancelFunc()
}

// Start runs the job
func (h HexBotJob) Start(done chan JobResult) {
	resp, err := api.CallHexbot(h.cancelContext, h.httpClient)
	done <- JobResult{ID: hexBotID, Result: resp, Err: err}
}

// NewHexBotJob creates a new job that returns a hexbot response
func NewHexBotJob(httpClient *http.Client) HexBotJob {
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	return HexBotJob{
		cancelContext,
		cancelFunc,
		httpClient,
	}
}
