package main

import (
	"fmt"
)

//State encapsulates the Brainfuck program's state
type state struct {
	data  []byte
	point int
}

func newState(size int) state {
	mem := make([]byte, size)
	return state{mem, 0}
}

func printState(state *state) {
	for _, i := range state.data {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func evalTok(cur byte, state *state) {
	if cur == '+' {
		state.data[state.point]++
	} else if cur == '-' {
		state.data[state.point]--
	} else if cur == '>' {
		state.point++
	} else if cur == '<' {
		if state.point > 0 {
			state.point--
		}
	} else if cur == '.' {
		fmt.Printf("%c", state.data[state.point])
	}
}

//Eval evaluates a tokenized brainfuck program
func Eval(ir *InternRep) {
	if ir.toks == nil {
		return
	}
	state := newState(30000) //according to Wikipedia, this is the "classic" size
	loopCount := -1
	for i := 0; i < len(ir.toks); i++ {
		cur := ir.toks[i]
		evalTok(cur, &state)
		if cur == '[' {
			loopCount++
			if state.data[state.point] == 0 {
				i = ir.loops[loopCount].end
			}
		} else if cur == ']' {
			i = ir.loops[loopCount].start - 1
			loopCount--
		}
	}
	fmt.Println()
}
