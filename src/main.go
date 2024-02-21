package main

import (
	"fmt"
	"hullo/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Hullo programming language!\n", user.Username)
	fmt.Printf("Executing the REPL!\n")
	repl.Start(os.Stdin, os.Stdout)
}
