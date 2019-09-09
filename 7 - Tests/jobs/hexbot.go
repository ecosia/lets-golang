package jobs

import (
	"context"

	"github.com/dohe/lets-golang/api"
	"github.com/dohe/lets-golang/common"
)

const (
	hexBotID = "hexbot"
)

// HexBotJob is the base type for this job
type HexBotJob struct {
	cancelContext context.Context
	cancelFunc    context.CancelFunc
	httpClient    common.HTTPClient
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
func NewHexBotJob(httpClient common.HTTPClient) HexBotJob {
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	return HexBotJob{
		cancelContext,
		cancelFunc,
		httpClient,
	}
}
