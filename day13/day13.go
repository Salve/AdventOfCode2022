package day13

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"sort"
	"strings"
)

const day = 13

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

const lb, rb = '[' - '0', ']' - '0'

func part1() {
	result := 0
	pairs := strings.Split(strings.ReplaceAll(string(input), "10", string('9'+1)), "\n\n")
	for i, pair := range pairs {
		s := strings.Split(pair, "\n")
		if rightOrder(s[0], s[1]) {
			result += i + 1
		}
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	dividers := elfPackets{"[[2]]", "[[6]]"}
	packets := append(elfPackets(strings.Fields(strings.ReplaceAll(string(input), "10", string('9'+1)))), dividers...)
	sort.Sort(packets)
	result := 1
	for i, p := range packets {
		if p == dividers[0] || p == dividers[1] {
			result *= i + 1
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

func rightOrder(a, b string) bool {
	left, right := a[0]-'0', b[0]-'0'
	if left == right {
		return rightOrder(a[1:], b[1:])
	}
	if left == rb {
		return true
	}
	if right == rb {
		return false
	}
	if left <= 10 && right <= 10 {
		return left < right
	}
	if left == lb {
		return rightOrder(a, "["+b[0:1]+"]"+b[1:])
	}
	if right == lb {
		return rightOrder("["+a[0:1]+"]"+a[1:], b)
	}
	panic(fmt.Sprintf("%s\n%s\n", a, b))
}

type elfPackets []string

func (x elfPackets) Len() int           { return len(x) }
func (x elfPackets) Less(i, j int) bool { return rightOrder(x[i], x[j]) }
func (x elfPackets) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
