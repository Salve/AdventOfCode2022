package day6

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
)

const day = 6

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	var result int
	for i := 4; i < len(input); i++ {
		if unique(input[i-4 : i]...) {
			result = i
			break
		}
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var result int
	for i := 14; i < len(input); i++ {
		if unique(input[i-14 : i]...) {
			result = i
			break
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

func unique(bs ...byte) bool {
	m := make(map[byte]struct{}, len(bs))
	for _, b := range bs {
		m[b] = struct{}{}
	}
	return len(m) == len(bs)
}
