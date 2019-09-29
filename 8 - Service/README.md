# 8. Using the results

## Steps

* Create a server
  * Create a `server` module & package
  * Create a `Serve` function, that starts a http server, serving `0.0.0.0:9876/`
  * Create a handleFunc that calls `colorpoint`'s `Get` and serializes the result to JSON
* Start the server in your `main.go` calling `server`'s `Serve`

## New concepts

| Concept | Go | Python |
|---|---|---|
| HTTP server | versatile [http server](https://golang.org/pkg/net/http/#Handler), including routing and building blocks for middlewares | basic [http server](https://docs.python.org/3/library/http.server.html) ⚠️ not recommended for production |
