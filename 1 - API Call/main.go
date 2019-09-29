package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

func main() { // PYTHON: if __name__ == "__main__"
	resp, err := api.CallHexbot()  // PYTHON: Error handling is done with an error object not a raise/except logic
	fmt.Printf("Error: %v\n", err) // PYTHON: print(f'Error: {err}')
	fmt.Printf("Response: %+v\n", resp)
}
