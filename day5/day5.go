package day5

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"strconv"
	"strings"
)

const day = 5

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
	lines := inputs.Lines(input)
	gap := findGap(lines)
	s := parseStacks(lines, gap)

	for _, line := range lines[gap+1:] {
		split := strings.Split(line, " ")
		crates, from, to := atoi(split[1]), atoi(split[3])-1, atoi(split[5])-1
		for i := 0; i < crates; i++ {
			s[to].Push(s[from].Pop(1))
		}
	}

	sb := strings.Builder{}
	for _, v := range s {
		sb.WriteByte(v.Pop(1)[0])
	}
	result := sb.String()
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	lines := inputs.Lines(input)
	gap := findGap(lines)
	s := parseStacks(lines, gap)

	for _, line := range lines[gap+1:] {
		split := strings.Split(line, " ")
		crates, from, to := atoi(split[1]), atoi(split[3])-1, atoi(split[5])-1
		s[to].Push(s[from].Pop(crates))
	}

	sb := strings.Builder{}
	for _, v := range s {
		sb.WriteByte(v.Pop(1)[0])
	}
	result := sb.String()
	fmt.Printf("Part 2: %v\n", result)
}

func (s *stack) Push(v []uint8) {
	*s = append(*s, v...)
}

func (s *stack) Pop(n int) []uint8 {
	l := len(*s)
	p := (*s)[l-n : l]
	*s = (*s)[:l-n]
	return p
}

func findGap(lines []string) int {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	panic("never")
}

type stack []uint8

func parseStacks(lines []string, gap int) []stack {
	c := (len(lines[0]) + 1) / 4
	s := make([]stack, c)
	for i := gap - 2; i >= 0; i-- {
		sn := 0
		for j := 1; j < c*4; j += 4 {
			if v := lines[i][j]; v != ' ' {
				s[sn] = append(s[sn], v)
			}
			sn++
		}
	}
	return s
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
