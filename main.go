package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/xxnmxx/gotax/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Gotax prompt!\n",user.Username)
	repl.Start(os.Stdin,os.Stdout)
}
