package day1

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"sort"
	"strconv"
	"strings"
)

const day = 1

var input = inputs.Input(day)

func init() {
	registry.Register(day, Run)
}

var ssums []int

func Run() {
	ssums = sortedSums()
	part1()
	part2()
}

func part1() {
	result := ssums[len(ssums)-1]
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	var result int
	for i := len(ssums); i > len(ssums)-3; i-- {
		result += ssums[i-1]
	}
	fmt.Printf("Part 2: %v\n", result)
}

func sortedSums() []int {
	var sums []int
	var sum int
	for _, line := range strings.Split(string(input)+"\n", "\n") {
		if line == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		v, _ := strconv.Atoi(line)
		sum += v
	}
	sort.Ints(sums)
	return sums
}
