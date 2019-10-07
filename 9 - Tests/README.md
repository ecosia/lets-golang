# 9. Tests

## Steps

* Use an interface for http clients 
  * Create a module called `type_definitions.go` in a module & package called `common`
  * Create a `HTTPClient` interface defined by a single method:
    * Do(req *http.Request) (*http.Response, error) (note that Go's standard `http.Client` implements this interface)
  * Change `colorpoint.Get` to accept `common.HTTPClient`
  * Do the same for the jobs and APIs
* Write tests for `colorpoint.go`
  * Create `colorpoint_test.go` right next to `colorpoint.go` (and make it part of the same package)
  * Create a `ClientMock` that implements `common.HTTPClient`
    * Make its struct have a reference to a `hexBody`, `vexBody` and `err` to return these based on the url
    * Hint: to create a `http.Response` `Body` from a string `s` you can write: `ioutil.NopCloser(strings.NewReader(s))`
  * Write a table test for `Get` for following cases:
    * Empty response
    * Empty json
    * Hex only
    * Vex only
    * Default (both Hex and Vex)
    * Error when calling
  * You can use the following mock bodies:
    * Hex:
      ```json
      {
        "colors": [
          { "value": "#FFAA00" }
        ]
      }
      ```
    * Vex:
      ```json
      {
        "vectors":
          [
            { 
              "a": { "x":1, "y":1 },
              "b": { "x":2, "y":2 }
            }
          ]
      }
      ```

## New concepts

| Concept | Go | Python |
|---|---|---|
| Mocking | no built-in support, relies heavily on [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) | `unittest.mock` [🔗](https://docs.python.org/3/library/unittest.mock.html) |
| Unit tests | `testing` [🔗](https://golang.org/pkg/testing/#pkg-overview) | `unittest` [🔗](https://docs.python.org/3/library/unittest.html#basic-example) |
| Parametrized/table tests | `testing.T.Run` [🔗](https://golang.org/pkg/testing/#T.Run) | `unittest.TestCase.subTest` [🔗](https://docs.python.org/3/library/unittest.html#distinguishing-test-iterations-using-subtests) (⚠️ since 3.4)
