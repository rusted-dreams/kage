package builtins

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "cd: missing argument")
		return
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "cd: too many arguments")
		return
	}
	dir := args[0]
	if err := os.Chdir(dir); err !=nil {
		fmt.Printf("cd: %v: No such file or directory\n", dir)
	}

}
