package main

import(
	"fmt"
	"os"
	"os/user"
	"hourglass.com/hourglass/src/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the Hourglass REPL!\n", user.Username)
	fmt.Printf("You can start writing code in Hourglass:\n")
	repl.Start(os.Stdin, os.Stdout)
}
