package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		// read user input:
		commandFullString, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading user input: %v\n", err)
			os.Exit(1)
		}
		commandFullString = strings.TrimSpace(commandFullString)
		tokens := strings.Split(commandFullString, " ")
		command := tokens[0]

		// process commands:
		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			if len(commandFullString) <= len(command) {
				fmt.Println()
				continue
			}
			fmt.Println(commandFullString[len(command)+1:])
		case "type":
			if len(tokens) < 2 {
				fmt.Fprintln(os.Stderr, "type: missing argument")
				continue
			}
			args := tokens[1]
			typeCommand(args)
		case "":
			continue
		default:
			execCommand(command, tokens[1:])
		}
	}
}

// execute external executables:
func execCommand(command string, args []string) {
	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%v: command not found\n", command)
		return
	}

	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
