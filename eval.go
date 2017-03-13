package main

import (
	"fmt"
)

const memsize = 1024

type state struct {
	data  []byte
	point int
}

func newState() state {
	mem := make([]byte, memsize)
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
		state.point--
	} else if cur == '.' {
		fmt.Printf("%c", state.data[state.point])
	} else if cur == ',' {
		state.data[state.point] = cur
	}
}

func loop(i int, toks []byte, state *state) int {
	if state.data[state.point] == 0 {
		loopCount := -1
		for {
			tmp := toks[i]
			if tmp == '[' {
				loopCount++
			} else if tmp == ']' && loopCount == 0 {
				break
			} else if tmp == ']' {
				loopCount--
			}
			i++
		}
	}
	return i
}

func unloop(i int, toks []byte, state *state) int {
	if state.data[state.point] != 0 {
		loopCount := -1
		for {
			tmp := toks[i]
			if tmp == ']' {
				loopCount++
			} else if tmp == '[' && loopCount == 0 {
				break
			} else if tmp == '[' {
				loopCount--
			}
			i--
		}
	}
	return i
}

//Eval evaluates a tokenized brainfuck program
func Eval(toks []byte) {
	if toks == nil {
		return
	}
	state := newState()
	for i := 0; i < len(toks); i++ {
		cur := toks[i]
		if cur == '[' {
			i = loop(i, toks, &state)
		} else if cur == ']' {
			i = unloop(i, toks, &state)
		} else {
			evalTok(cur, &state)
		}
	}
}
