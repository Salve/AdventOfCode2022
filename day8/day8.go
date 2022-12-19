package day8

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"strconv"
)

const day = 8

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	// TODO: try this with image.Point and multiplying by directions
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	lines := inputs.Lines(input)
	seen := map[string]struct{}{}
	lastLine := len(lines) - 1
	lastCol := len(lines[0]) - 1
	// left
	for i := 0; i <= lastLine; i++ {
		max := -1
		for j := 0; j <= lastCol; j++ {
			v := atoi(string(lines[i][j]))
			if v > max {
				max = v
				seen[fmt.Sprintf("%d,%d", i, j)] = struct{}{}
			}
		}
	}

	// right
	for i := 0; i <= lastLine; i++ {
		max := -1
		for j := lastCol; j >= 0; j-- {
			v := atoi(string(lines[i][j]))
			if v > max {
				max = v
				seen[fmt.Sprintf("%d,%d", i, j)] = struct{}{}
			}
		}
	}

	// top
	for j := 0; j <= lastCol; j++ {
		max := -1
		for i := 0; i <= lastLine; i++ {
			v := atoi(string(lines[i][j]))
			if v > max {
				max = v
				seen[fmt.Sprintf("%d,%d", i, j)] = struct{}{}
			}
		}
	}

	// bottom
	for j := 0; j <= lastCol; j++ {
		max := -1
		for i := lastLine; i >= 0; i-- {
			v := atoi(string(lines[i][j]))
			if v > max {
				max = v
				seen[fmt.Sprintf("%d,%d", i, j)] = struct{}{}
			}
		}
	}

	result := len(seen)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	m := inputs.Lines(input)
	result := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			p := point{i, j}
			sum := treesSeen(m, p, left) * treesSeen(m, p, right) * treesSeen(m, p, up) * treesSeen(m, p, down)
			if sum > result {
				result = sum
			}
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}

func treesSeen(m []string, pos, dir point) (total int) {
	start := atoi(string(m[pos.i][pos.j]))
	for {
		pos.i += dir.i
		pos.j += dir.j
		if !valid(m, pos) {
			return total
		}
		total += 1
		if atoi(string(m[pos.i][pos.j])) >= start {
			break
		}
	}
	return total
}

func valid(m []string, p point) bool {
	return p.i >= 0 && p.i < len(m) && p.j >= 0 && p.j < len(m[0])
}

var (
	left  = point{i: 0, j: -1}
	right = point{i: 0, j: 1}
	up    = point{i: -1, j: 0}
	down  = point{i: 1, j: 0}
)

type point struct {
	i, j int
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
