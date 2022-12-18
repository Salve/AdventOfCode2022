package day3

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
)

const day = 3

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
	for _, line := range inputs.Lines(input) {
		half := len(line) / 2
		result += common(line[:half], line[half:])
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var result int
	lines := inputs.Lines(input)
	for i := 0; i < len(lines); i += 3 {
		result += common(lines[i], lines[i+1], lines[i+2])
	}
	fmt.Printf("Part 2: %v\n", result)
}

func common(lines ...string) int {
	var itemCount [53]int
	elfID, elfAll := 1, 0
	for _, line := range lines {
		elfID <<= 1
		elfAll |= elfID
		for _, r := range line {
			itemCount[priority(r)] |= elfID
		}
	}
	for p, sum := range itemCount {
		if sum == elfAll {
			return p
		}
	}
	panic("never")
}

func priority(r rune) int {
	if r > 'Z' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}
