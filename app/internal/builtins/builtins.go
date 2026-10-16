package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var BuiltInCommandList = make(map[string]func(args []string))

func init() {
	BuiltInCommandList["echo"] = Echo
	BuiltInCommandList["exit"] = Exit
	BuiltInCommandList["type"] = Type
}

func IsBuiltInCommand(command string) bool {
	_, ok := BuiltInCommandList[command]
	return ok
}

func RunBuiltIn(command string, args []string) {
	ok := IsBuiltInCommand(command)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s is not a shell builtin\n", command)
		return
	}
	cmd := BuiltInCommandList[command]
	cmd(args)
}

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func Exit(args []string) {
	os.Exit(0)
}

func Type(args []string) {

	for _, cmd := range args {
		if _, ok := BuiltInCommandList[cmd]; ok {
			fmt.Printf("%v: is a shell builtin\n", cmd)
			continue
		}
		
		// check if the command is a sys executable
		// if input is like PATH="/usr/bin:/usr/local/bin:$PATH" ./your_program.sh
		// then recursively checks all dirs. for the executable.
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("%v: not found\n", cmd)
			continue
		}
		fmt.Printf("%v is %v\n", args, path)
	}

}
