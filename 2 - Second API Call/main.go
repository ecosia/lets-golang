package main

import (
	"fmt"

	"github.com/dohe/lets-golang/api"
)

func main() {
	respHex, err := api.CallHexbot()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Response: %+v\n", respHex)

	respVex, err := api.CallVexbot()
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Response: %+v\n", respVex)
}
