package jobs

type JobResult struct {
	ID     string
	Result interface{}
	Err    error
}

type Job interface { // PYTHON: You'd do something similar with (abstract) classes and inheritance
	ID() string
	Start(done chan JobResult)
	Cancel()
}
