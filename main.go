package main

import (
	"fmt"
	"geomys/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the geomys programming language.\n", usr.Username)
	fmt.Printf("Start typing code!\n")
	repl.Start(os.Stdin, os.Stdout)
}
