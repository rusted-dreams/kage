package builtins

import (
	"fmt"
	"os"
)

func Pwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "pwd:", err)
		return
	}
	
	fmt.Println(dir)
}
