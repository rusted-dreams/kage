package main

import (
	"fmt"
	"os/exec"
	"slices"
)

// builtin commands arr
var builtIns = []string{
	"echo",
	"exit",
	"type",
}

func typeCommand(args string) {
	if slices.Contains(builtIns, args) {
		fmt.Printf("%v: is a shell builtin\n", args)
		return
	}

	// check if the command is a sys executable
	// if input is like PATH="/usr/bin:/usr/local/bin:$PATH" ./your_program.sh
	// then recursively checks all dirs. for the executable.
	path, err := exec.LookPath(args)
	if err != nil {
		fmt.Printf("%v: not found\n", args)
		return
	}
	fmt.Printf("%v is %v\n", args, path)
}
