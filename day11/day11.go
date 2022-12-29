package day11

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"sort"
	"strconv"
	"strings"
)

const day = 11

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
	ms := monkeys(string(input))
	business := make([]int, len(ms))
	for i := 0; i < 20; i++ {
		for m := range ms {
			for len(ms[m].items) > 0 {
				business[m]++
				item := ms[m].items[0]
				ms[m].items = ms[m].items[1:]
				item = ms[m].inspect(item)
				item /= 3
				next := ms[m].test(item)
				ms[next].items = append(ms[next].items, item)
			}
		}
	}
	sort.Ints(business)
	result := business[len(business)-2] * business[len(business)-1]
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	ms := monkeys(string(input))
	business := make([]int, len(ms))
	productDivisors := 1
	for _, m := range ms {
		productDivisors *= m.divisor
	}
	for i := 0; i < 10000; i++ {
		for m := range ms {
			for len(ms[m].items) > 0 {
				business[m]++
				item := ms[m].items[0]
				ms[m].items = ms[m].items[1:]
				item = ms[m].inspect(item)
				next := ms[m].test(item)
				ms[next].items = append(ms[next].items, item%productDivisors) // did not know the modulo trick
			}
		}
	}
	sort.Ints(business)
	result := business[len(business)-2] * business[len(business)-1]
	fmt.Printf("Part 2: %v\n", result)
}

type monkey struct {
	items   []int
	inspect func(int) int
	test    func(int) int
	divisor int
}

func monkeys(descs string) []monkey {
	var o []monkey
	for _, desc := range strings.Split(descs, "\n\n") {
		o = append(o, newMonkey(desc))
	}
	return o
}

func newMonkey(desc string) monkey {
	lines := strings.Split(desc, "\n")
	t, divisor := test(lines[3], lines[4], lines[5])
	return monkey{
		items:   items(lines[1]),
		inspect: inspect(lines[2]),
		test:    t,
		divisor: divisor,
	}
}

func test(desc ...string) (func(int) int, int) {
	var vs []int
	for _, line := range desc {
		s := strings.Split(line, " ")
		vs = append(vs, atoi(s[len(s)-1]))
	}
	return func(v int) int {
		if v%vs[0] == 0 {
			return vs[1]
		}
		return vs[2]
	}, vs[0]
}

func inspect(desc string) func(int) int {
	s := strings.Split(desc, " ")
	op, mod := s[len(s)-2], s[len(s)-1]
	return func(v int) int {
		var m int
		if mod == "old" {
			m = v
		} else {
			m = atoi(mod)
		}
		switch op {
		case "+":
			return v + m
		case "*":
			return v * m
		default:
			panic("never")
		}
	}
}

func items(desc string) []int {
	s := strings.Split(desc, ": ")
	vals := strings.Split(s[1], ", ")
	var o []int
	for _, v := range vals {
		o = append(o, atoi(v))
	}
	return o
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
