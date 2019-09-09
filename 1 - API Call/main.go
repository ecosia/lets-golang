package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

func main() { // PYTHON: if __name__ == "__main__"
	resp, err := api.CallHexbot()
	fmt.Printf("Error: %v\n", err) // PYTHON: print(f'Error: {err}')
	fmt.Printf("Error: %+v\n", resp)
}
