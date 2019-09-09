package jobs

type JobResult struct {
	ID      string
	Result  interface{}
	Err     error
}

type Job interface {
	ID() string
	Start(done chan JobResult)
	Cancel()
}