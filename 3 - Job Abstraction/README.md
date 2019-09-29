# 3. Job Abstraction

## Steps

* Write base job abstraction
  * Create a new folder & package `jobs`
  * Write a `JobResult` struct, with:
    * the job's `ID` (`string`)
    * the job's `Result` (we don't know the type beforehand, so `interface{}`)
    * a potential error `Err` (`error`)
  * Write a `Job` interface, defined by:
    * `ID()`
    * `Start(done chan JobResult)`
    * `Cancel()`
* Write a hexbot job (`jobs/hexbot.go`)
  * Follow the base `Job` interface (`Cancel` will be a no-op for now)
  * In `Start`, call `api.CallHexbot()` and write the `JobResult` to the provided channel
* Write a vexbot job (`jobs/vexbot.go`)
* Start both jobs in the main
  * Use the `go` keyword to start a new goroutine
  * Create a channel and pass it to the jobs
  * Wait on the channel until both jobs have send a message
  * Use a type assertion switch/case to print what job the received `JobResult` came from

## New concepts

| Concept | Go | Python |
|---|---|---|
| Logging | `log` [🔗](https://golang.org/pkg/log/) | `logging` [🔗](https://docs.python.org/3/library/logging.html) |
| Synchronizing threaded execution | The `go` keyword and channels [🔗](https://tour.golang.org/concurrency/2) | Possible to implement with a combination of [threads](https://docs.python.org/3.7/library/threading.html) and a [queue](https://docs.python.org/3/library/queue.html) |
| Loops | only `for` [🔗](https://tour.golang.org/flowcontrol/1) | `for` and `while` [🔗](https://www.learnpython.org/en/Loops) |
| Interfaces | No inheritance, instead [interface definitions](https://tour.golang.org/methods/10) and [struct embedding](https://golang.org/doc/effective_go.html#embedding) | [Classical inheritance](https://www.w3schools.com/python/python_inheritance.asp) |
| Methods | Functions with [receivers](https://tour.golang.org/methods/1) | As part of the [class definition](https://docs.python.org/3/tutorial/classes.html#class-objects) |
| Type assertion | `.(type)` [🔗](https://tour.golang.org/methods/15) | `isinstance` [🔗](https://docs.python.org/3.7/library/functions.html#isinstance) |
