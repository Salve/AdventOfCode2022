package day9

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
)

const day = 9

var input []byte
var dirs = map[string]image.Point{"R": {X: 1, Y: 0}, "L": {X: -1, Y: 0}, "U": {X: 0, Y: -1}, "D": {X: 0, Y: 1}}

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	h, t := image.Point{}, image.Point{}
	visited := map[image.Point]struct{}{t: {}}
	for _, line := range inputs.Lines(input) {
		s := strings.Split(line, " ")
		for i := 0; i < atoi(s[1]); i++ {
			h = h.Add(dirs[s[0]])
			t = newTailPos(t, h)
			visited[t] = struct{}{}
		}
	}
	result := len(visited)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	pos := make([]image.Point, 10)
	visited := map[image.Point]struct{}{pos[9]: {}}
	for _, line := range inputs.Lines(input) {
		s := strings.Split(line, " ")
		for i := 0; i < atoi(s[1]); i++ {
			pos[0] = pos[0].Add(dirs[s[0]])
			for j := 1; j <= 9; j++ {
				pos[j] = newTailPos(pos[j], pos[j-1])
			}
			visited[pos[9]] = struct{}{}
		}
	}
	result := len(visited)
	fmt.Printf("Part 2: %v\n", result)
}

func newTailPos(t, h image.Point) image.Point {
	diff := h.Sub(t)
	touching := abs(diff.X) <= 1 && abs(diff.Y) <= 1
	if diff.X > 1 || !touching && diff.X > 0 {
		t.X += 1
	}
	if diff.X < -1 || !touching && diff.X < 0 {
		t.X -= 1
	}
	if diff.Y > 1 || !touching && diff.Y > 0 {
		t.Y += 1
	}
	if diff.Y < -1 || !touching && diff.Y < 0 {
		t.Y -= 1
	}
	return t
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
