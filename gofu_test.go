package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const scriptDir = "./test_scripts"

func TestSimple(t *testing.T) {
	files, err := ioutil.ReadDir(scriptDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		file, err := ioutil.ReadFile(scriptDir + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		interpret(file)
	}
}
