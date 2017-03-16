package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 { //if there's no program to run, don't run :)
		return
	}
	//can't think of a way to chunk up a bf program -- read it all in!
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	interpret(file)
}
