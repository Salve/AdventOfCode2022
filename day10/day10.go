package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
)

const day = 10

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
	x := 1
	during := []int{x}
	for _, line := range inputs.Lines(input) {
		s := strings.Split(line, " ")
		switch s[0] {
		case "noop":
			during = append(during, x)
		case "addx":
			during = append(during, x, x)
			x += atoi(s[1])
		}
	}
	var result int
	for i := 20; i <= 220; i += 40 {
		result += during[i] * i
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	x := 1
	during := []int{x}
	for _, line := range inputs.Lines(input) {
		s := strings.Split(line, " ")
		switch s[0] {
		case "noop":
			during = append(during, x)
		case "addx":
			during = append(during, x, x)
			x += atoi(s[1])
		}
	}
	var sb strings.Builder
	for row := 0; row < 6; row++ {
		for pos := 0; pos < 40; pos++ {
			cycle := 40*row + pos + 1
			sprite := during[cycle]
			if sprite-1 == pos || sprite == pos || sprite+1 == pos {
				sb.WriteString("##")
			} else {
				sb.WriteString("  ")
			}
		}
		sb.WriteString("\n")
	}
	result := sb.String()
	fmt.Printf("Part 2: \n%v\n", result)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
