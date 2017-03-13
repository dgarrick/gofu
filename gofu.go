package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 { //if there's no program to run, don't run :)
		return
	}
	Eval(Parse(os.Args[1]))
}
