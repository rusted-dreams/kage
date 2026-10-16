package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rusted-dreams/kage/app/internal/builtins"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		// read user input:
		commandFullString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading user input: %v\n", err)
			os.Exit(1)
		}
		commandFullString = strings.TrimSpace(commandFullString)
		tokens := strings.Split(commandFullString, " ")
		command := tokens[0]
		run(command, tokens[1:])
	}
}

func run(command string, args []string) {
	if ok := builtins.IsBuiltInCommand(command); ok {
		builtins.RunBuiltIn(command, args)
		return
	}
	execExtCommand(command, args)
}

// execute external executables:
func execExtCommand(command string, args []string) {
	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%v: command not found\n", command)
		return
	}

	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin // var builtIns = []string{"echo", "exit", "type"}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
