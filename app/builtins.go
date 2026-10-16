package main

import (
	"fmt"
	"slices"
)

// builtin commands arr
var builtIns = []string{
	"echo",
	"exit",
	"type",
}

func typeCommand(args string) {
	// for _, builtIn := range builtIns {
	// 	if args == builtIn {
	// 		fmt.Printf("%v: is a shell builtin\n", args)
	// 		return
	// 	}
	// }

	if slices.Contains(builtIns, args) {
		fmt.Printf("%v: is a shell builtin\n", args)
		return
	}
	fmt.Printf("%v: not found\n", args)
}
