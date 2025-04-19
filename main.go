package main

import (
	"github.com/chzyer/readline"
)

func main() {
	var err error
	rl, err = readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	calculate()

}
