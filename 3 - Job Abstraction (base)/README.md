# 3. Job Abstraction (base)

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

## New concepts

| Concept | Go | Python |
|---|---|---|
| Interfaces | No inheritance, instead [interface definitions](https://tour.golang.org/methods/10) and [struct embedding](https://golang.org/doc/effective_go.html#embedding) | [Classical inheritance](https://www.w3schools.com/python/python_inheritance.asp) |
| Methods | Functions with [receivers](https://tour.golang.org/methods/1) | As part of the [class definition](https://docs.python.org/3/tutorial/classes.html#class-objects) |
