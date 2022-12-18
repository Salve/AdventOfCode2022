package day2

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"strings"
)

const day = 2

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
		a, b := splitLine(line)
		they, you := signFromSymbol(a), signFromSymbol(b)
		result += fight(they, you) + int(you)
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var result int
	for _, line := range inputs.Lines(input) {
		a, b := splitLine(line)
		they := signFromSymbol(a)
		you := signFromSignal(they, b)
		result += fight(they, you) + int(you)
	}
	fmt.Printf("Part 2: %v\n", result)
}

func splitLine(line string) (they, you string) {
	s := strings.Split(line, " ")
	return s[0], s[1]
}

type sign int

const (
	rock sign = iota + 1
	paper
	scissor
)

var signals = map[string]int{"X": 0, "Y": 3, "Z": 6}

func signFromSignal(they sign, signal string) sign {
	for _, you := range []sign{rock, paper, scissor} {
		result := fight(they, you)
		if result == signals[signal] {
			return you
		}
	}
	panic("never")
}

func signFromSymbol(s string) sign {
	switch s {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissor
	}
	panic("never")
}

func fight(they, you sign) int {
	if they == you {
		return 3
	}
	if (you == rock && they == scissor) || (you == paper && they == rock) || (you == scissor && they == paper) {
		return 6
	}
	return 0
}
