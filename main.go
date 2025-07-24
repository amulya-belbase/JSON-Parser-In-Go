package main

import (
	"fmt"
	"json_parser/resolver"
)

func main() {
	// send action & processId to resolver
	resolverData := resolver.ResolveWithActionProcessID("sub", "sub_1")

	// get the resolver data back
	fmt.Println("Resolved data: ", resolverData)
}
