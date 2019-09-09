package jobs

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

const (
	hexBotID = "hexbot"
)

// HexBotJob is the base type for this job
type HexBotJob struct{}

// ID returns the job's ID
func (h HexBotJob) ID() string {
	return hexBotID
}

// Cancel isn't implemented yet
func (h HexBotJob) Cancel() {
	fmt.Println("TODO: implement cancel") //TODO: implement cancel
}

// Start runs the job
func (h HexBotJob) Start(done chan JobResult) {
	resp, err := api.CallHexbot()
	done <- JobResult{ID: hexBotID, Result: resp, Err: err}
}

// NewHexBotJob creates a new job that returns a hexbot response
func NewHexBotJob() HexBotJob {
	return HexBotJob{}
}
