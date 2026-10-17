package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/rusted-dreams/kage/app/internal/builtins"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		tokens, err := getInput(reader)
		if err != nil {
			fmt.Printf("err reading user input: %s\n", err)
		}
		if len(tokens) > 0 {
			run(tokens[0], tokens[1:])
		}
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

func getInput(reader *bufio.Reader) ([]string, error) {
	var tokens []string
	var token strings.Builder

	tokenStarted := false

	saveToken := func() {
		if tokenStarted {
			tokens = append(tokens, token.String())
			token.Reset()
			tokenStarted = false
		}
	}

	inSingleQuote := false
	inDoubleQuote := false
	isEscapeChar := false

	for {
		if inSingleQuote || inDoubleQuote {
			fmt.Print("> ")
		}
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		for _, char := range line {
			if isEscapeChar {
				token.WriteRune(rune(char))
				isEscapeChar = false
				continue
			}
			if char == '\\' && !inSingleQuote {
				isEscapeChar = true
				continue
			}
			if char == '\'' && !inDoubleQuote {
				inSingleQuote = !inSingleQuote
				tokenStarted = true
			} else if char == '"' && !inSingleQuote {
				inDoubleQuote = !inDoubleQuote
				tokenStarted = true
			} else if inSingleQuote {
				token.WriteRune(rune(char))
				tokenStarted = true
			} else if inDoubleQuote {
				token.WriteRune(rune(char))
				tokenStarted = true
			} else if char == ' ' || char == '\t' || char == '\n' {
				saveToken()
			} else {
				tokenStarted = true
				token.WriteRune(rune(char))
			}
		}

		if inSingleQuote {
			continue
		}
		if inDoubleQuote {
			continue
		}

		return tokens, nil
	}
}
