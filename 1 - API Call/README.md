# 1. API Call

## Steps

* Create a main entry point
  * Create a file called `main.go`
  * Make it part of the `main` package
  * Create a `main` function
* Write code to call first API
  * Create folder & package `api`
  * Create a module called `hexbot.go`
  * Call Github's [hexbot API](https://api.noopschallenge.com/hexbot) in a `CallHexbot() (*HexbotResponse, error)` function
  * Unmarshal JSON response to a `HexbotResponse` struct
* Run your code with `go run main.go`

## New concepts

| Concept | Go | Python |
|---|---|---|
| Module | `package` and `import` [🔗](https://tour.golang.org/basics/1) | modules [🔗](https://docs.python.org/3/tutorial/modules.html) |
| Main | `def main()` | `if __name__ == "__main__"` |
| Printing | `fmt.Println` [🔗](https://golang.org/pkg/fmt/#Print) | `print` [🔗](https://docs.python.org/3/library/functions.html#print) |
| String formatting | `fmt.Printf` [🔗](https://golang.org/pkg/fmt/#Printf) | `f'{var}'` [🔗](https://docs.python.org/3/tutorial/inputoutput.html#fancier-output-formatting) |
| Types | Go has [static types](https://go101.org/article/type-system-overview.html) with inferred types | Python has [type annotations](https://docs.python.org/3/library/typing.html) and duck typing |
| Structs | Go has [structs](https://tour.golang.org/moretypes/2) | In python you might use [data classes](https://docs.python.org/3/library/dataclasses.html) |
| Making HTTP calls | [http](https://golang.org/pkg/net/http/#Get) | [urllib](https://docs.python.org/3/library/urllib.request.html#urllib.request.urlopen) or [requests](https://2.python-requests.org/en/master/) |
| Error handling | Explicit [error object](https://blog.golang.org/error-handling-and-go) | `raise` and `except` [🔗](https://docs.python.org/3/tutorial/errors.html) |
| JSON parsing | Uses [json package](https://golang.org/pkg/encoding/json/#Unmarshal), trickier because of types | [json](https://docs.python.org/3/library/json.html#json.loads) |
| Pointers | Yup [🔗](https://tour.golang.org/moretypes/1) | Nope, but e.g. lists and dictionaries are passed by reference. Or not... tricky question [🔗](https://www.jeffknupp.com/blog/2012/11/13/is-python-callbyvalue-or-callbyreference-neither/) |
| Null type | `nil`, but only for pointers and some complex types, otherwise default values  [🔗](https://go101.org/article/nil.html) | `None` [🔗](https://docs.python.org/3/library/constants.html#None) |
