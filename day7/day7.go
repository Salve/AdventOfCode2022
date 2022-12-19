package day7

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"math"
	"strconv"
	"strings"
)

const day = 7

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
	sizes := dirSizes(inputs.Lines(input))
	result := 0
	for _, v := range sizes {
		if v <= 100_000 {
			result += v
		}
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	sizes := dirSizes(inputs.Lines(input))
	missingSpace := 30_000_000 - (70_000_000 - sizes["/"])
	result := math.MaxInt
	for _, v := range sizes {
		if v >= missingSpace && v < result {
			result = v
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

func dirSizes(lines []string) map[string]int {
	var stack []string
	sizes := make(map[string]int)
	for _, line := range lines {
		s := strings.Split(line, " ")
		switch s[0] {
		case "$":
			if s[1] != "cd" {
				continue
			}
			if s[2] == ".." {
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, strings.Join(append(stack, s[2]), "/"))
			continue
		case "dir":
			continue
		default:
			for _, dir := range stack {
				sizes[dir] += atoi(s[0])
			}
		}
	}
	return sizes
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
