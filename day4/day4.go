package day4

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"strconv"
	"strings"
)

const day = 4

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
		if contained(split(line)) {
			result++
		}
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var result int
	for _, line := range inputs.Lines(input) {
		if overlap(split(line)) {
			result++
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

func contained(astart, aend, bstart, bend int) bool {
	return (astart <= bstart && aend >= bend) || (bstart <= astart && bend >= aend)
}

func overlap(astart, aend, bstart, bend int) bool {
	return astart <= bend && aend >= bstart
}

func split(line string) (astart, aend, bstart, bend int) {
	split := strings.Split(line, ",")
	a := strings.Split(split[0], "-")
	b := strings.Split(split[1], "-")
	return atoi(a[0]), atoi(a[1]), atoi(b[0]), atoi(b[1])
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
