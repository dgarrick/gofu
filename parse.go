package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//InternRep is the internal representation of a brainfuck program
type InternRep struct {
	toks  []byte
	loops []Loop
}

//Loop represents the positions of a paired open and close bracket
type Loop struct {
	start int
	end   int
}

func newInternRep() *InternRep {
	var toks []byte
	var blocks []Loop
	return &InternRep{toks, blocks}
}

//Parse parses a filename to a slice of brainfuck tokens
func Parse(fname string) *InternRep {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	ir := newInternRep()
	var bracks []int
	isComm := false
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
			if (cur == '+' ||
				cur == '-' ||
				cur == '>' ||
				cur == '<' ||
				cur == '.' ||
				cur == ',') && !isComm {
				charNum++
				ir.toks = append(ir.toks, cur)
			} else if cur == '"' {
				isComm = !isComm
				break
			} else if cur == '[' {
				loopCount++
				charNum++
				ir.toks = append(ir.toks, cur)
				bracks = append(bracks, charNum)
			} else if cur == ']' {
				loopCount--
				charNum++
				ir.toks = append(ir.toks, cur)
				if ir.toks[bracks[len(bracks)-1]] == '[' {
					pop := bracks[len(bracks)-1]
					bracks = bracks[:len(bracks)-1]
					ir.loops = append(ir.loops, Loop{pop, charNum})
				} else {
					fmt.Println("Unbalanced brackets!")
					return nil
				}
			}
		}
	}
	for _, c := range ir.loops {
		fmt.Printf("%d\n", c)
	}
	if loopCount != 0 {
		fmt.Println("Unbalanced brackets!")
		return nil
	}
	return ir
}
