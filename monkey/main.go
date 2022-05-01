package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 1 {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) >= 2 {
		if len(os.Args) == 2 {
			repl.FileREP(os.Args[1], os.Stdout, make([]string, 0))
		} else {
			repl.FileREP(os.Args[1], os.Stdout, os.Args[2:])
		}
	}
}
