package main

import (
	"bufio"
	"fmt"
	"os"
)

//memory size defaults to 30kb
const memsize = 30000

type state struct {
	data   []byte
	point  int
	bracks []int
}

func newState() state {
	mem := make([]byte, memsize)
	var bracks []int
	return state{mem, 0, bracks}
}

func printState(state *state) {
	for _, i := range state.data {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func evalTok(cur byte, state *state, readIn *bufio.Reader) {
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
		tmp, err := readIn.ReadByte()
		if err != nil {
			fmt.Println(err)
		} else {
			state.data[state.point] = tmp
		}
	}
}

func loop(i int, cur byte, toks []byte, state *state) int {
	if state.data[state.point] == 0 {
		loopCount := 1
		for loopCount > 0 {
			i++
			cur = toks[i]
			if cur == '[' {
				loopCount++
			} else if cur == ']' {
				loopCount--
			}
		}
	} else {
		//push the bracket on the stack
		state.bracks = append(state.bracks, i)
	}
	return i
}

func unloop(i int, cur byte, state *state) int {
	if state.data[state.point] != 0 {
		return state.bracks[len(state.bracks)-1]
	}
	//pop the bracket stack and return the current index
	state.bracks = state.bracks[:len(state.bracks)-1]
	return i
}

func interpret(toks []byte) {
	state := newState()
	readIn := bufio.NewReader(os.Stdin)
	var cur byte
	for i := 0; i < len(toks); i++ {
		cur = toks[i]
		if cur == '[' {
			i = loop(i, cur, toks, &state)
		} else if cur == ']' {
			i = unloop(i, cur, &state)
		} else {
			evalTok(cur, &state, readIn)
		}
	}
}
