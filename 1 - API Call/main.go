package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

func main() {
	resp, err := api.CallHexbot()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Error: %+v\n", resp)
}
