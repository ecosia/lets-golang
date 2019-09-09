package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

func main() {
	respHex, err := api.CallHexbot()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Error: %+v\n", respHex)

	respVex, err := api.CallVexbot()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Error: %+v\n", respVex)
}
