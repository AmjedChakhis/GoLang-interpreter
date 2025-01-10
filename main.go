package main

import (
	"fmt"
	"os"
	"os/user"

	repl "github.com/AmjedChakhis/GoLang-interpreter/REPL"
)

func main() {
	currUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, welcome to GBI : Go-based-Interpreter ", currUser.Username)
	fmt.Printf("Start typing your code ...\n")

	repl.Start(os.Stdin, os.Stdout)
}
