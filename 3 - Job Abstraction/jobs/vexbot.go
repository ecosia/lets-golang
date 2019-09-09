package jobs

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

const (
	vexBotID = "vexbot"
)

// VexBotJob is the base type for this job
type VexBotJob struct{}

// ID returns the job's ID
func (v VexBotJob) ID() string {
	return vexBotID
}

// Cancel isn't implemented yet
func (v VexBotJob) Cancel() {
	fmt.Println("TODO: implement cancel") //TODO: implement cancel
}

// Start runs the job
func (v VexBotJob) Start(done chan JobResult) {
	resp, err := api.CallVexbot()
	done <- JobResult{ID: vexBotID, Result: resp, Err: err}
}

// NewVexBotJob creates a new job that returns a vexbot response
func NewVexBotJob() VexBotJob {
	return VexBotJob{}
}
