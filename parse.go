package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Parse parses a filename to a slice of brainfuck tokens
func Parse(fname string) []byte {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	var toks []byte
	loopCount := 0
	charNum := -1
	r := bufio.NewReader(file)
	for {
		if cur, err := r.ReadByte(); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Fatal error encountered for %c", cur)
				return nil
			}
		} else {
			if cur == '+' ||
				cur == '-' ||
				cur == '>' ||
				cur == '<' ||
				cur == '.' ||
				cur == ',' ||
				cur == '[' ||
				cur == ']' {
				charNum++
				toks = append(toks, cur)
			}
			if cur == '[' {
				loopCount++
			}
			if cur == ']' {
				loopCount--
			}
		}
	}
	if loopCount != 0 {
		fmt.Println("Unbalanced brackets!")
		return nil
	}
	return toks
}
