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
	state := newState(10) //according to Wikipedia, this is the "classic" size
	loopCount := -1
	loopIdx := -1
	var loopStack []int
	loopStack = append(loopStack, 0)
	for i := 0; i < len(ir.toks); i++ {
		cur := ir.toks[i]
		evalTok(cur, &state)
		if cur == '[' {
			loopCount++
			//FIXME we can't reliably reset the outermost loop
			if loopCount >= len(ir.loops) {
				loopCount = 0
			}
			loopIdx = loopCount
			//if this is our first time in the loop, add it to the stack
			if loopStack[len(loopStack)-1] != loopIdx {
				loopStack = append(loopStack, loopCount)
				fmt.Println(loopStack)
			}
			if state.data[state.point] == 0 {
				i = ir.loops[loopIdx].end - 1
			}
		} else if cur == ']' {
			//FIXME this is broken. We need to keep track of how many loops we've seen
			//and which index in the Loop we're currently in.
			if state.data[state.point] != 0 {
				i = ir.loops[loopIdx].start - 1
				loopCount--
			} else {
				loopStack = loopStack[:len(loopStack)-1]
				loopIdx = loopStack[len(loopStack)-1]
			}
		}
	}
	fmt.Println()
}
