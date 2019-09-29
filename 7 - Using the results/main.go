package main

import (
	"fmt"

	"github.com/dohe/lets-golang/colorpoint"
)

func main() {
	res := colorpoint.Get()
	fmt.Printf("Got combined result: %+v\n", res)
}
