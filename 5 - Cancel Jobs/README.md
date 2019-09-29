# 5. Cancel jobs

## Steps

* Make APIs cancellable
  * Make the `Call` functions accept a `context.Context` and a `*http.Client`
  * Create a request with the given context and make the http call with the provided client
* Make jobs provide the context and client APIs now expect
  * In your `New` function expect a http client to pass on to the API
  * Create a cancel context. Keep a reference to it in your job's struct, so you can interact with it on `Cancel`
* Implement cancelling on timeout
  * Add a timeout case, iterating over all jobs and calling their `Cancel` method
  * Adjust your waiting logic to account for a potential timeout

## New concepts

| Concept | Go | Python |
|---|---|---|
| Connection reuse | Sharing a `http.Client` pointer [🔗](https://golang.org/pkg/net/http/#Client.Do) (⚠️ make sure to read and close the body) | With `urllib3.HTTPConnectionPool` (also used by `requests`'s `Session`) [🔗](https://urllib3.readthedocs.io/en/1.4/pools.html) |
| Canceling requests | [cancel context](https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/) | not easily possible with urllib/requests |
