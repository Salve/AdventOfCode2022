package day14

import (
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"image"
	"regexp"
	"strconv"
)

const day = 14

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
	c := rockScanner(input)
	maxY := 0
	for p := range c {
		if y := p.Y; y > maxY {
			maxY = y
		}
	}
	result := 0
	for c.sandSettled(image.Point{500, 0}, maxY) {
		result++
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	c := rockScanner(input)
	maxY := 0
	for p := range c {
		if y := p.Y; y > maxY {
			maxY = y
		}
	}
	result := 0
	for !c.sandBlocked(image.Point{500, 0}, maxY) {
		result++
	}
	fmt.Printf("Part 2: %v\n", result+1)
}

type cave map[image.Point]struct{}

func (c cave) sandBlocked(sand image.Point, maxY int) bool {
	if sand.Y >= maxY+1 {
		// sand is on the floor
		c[sand] = struct{}{}
		return false
	}
	for _, dir := range []image.Point{{0, 1}, {-1, 1}, {1, 1}} {
		try := sand.Add(dir)
		if _, blocked := c[try]; !blocked {
			return c.sandBlocked(try, maxY)
		}
	}
	// all directions blocked
	if sand.Y == 0 {
		return true
	}
	// sand has settled
	c[sand] = struct{}{}
	return false
}

func (c cave) sandSettled(sand image.Point, maxY int) bool {
	if sand.Y > maxY {
		return false
	}
	for _, dir := range []image.Point{{0, 1}, {-1, 1}, {1, 1}} {
		try := sand.Add(dir)
		if _, blocked := c[try]; !blocked {
			return c.sandSettled(try, maxY)
		}
	}
	// all directions blocked, sand has settled
	c[sand] = struct{}{}
	return true
}

func rockScanner(scan []byte) cave {
	c := make(cave)
	for _, line := range inputs.Lines(scan) {
		for _, p := range points(paths(line)) {
			c[p] = struct{}{}
		}
	}
	return c
}

func points(path []image.Point) []image.Point {
	next := 1
	x, y := path[0].X, path[0].Y
	ps := []image.Point{{x, y}}
	for {
		if path[next].X > x {
			x++
		}
		if path[next].X < x {
			x--
		}
		if path[next].Y > y {
			y++
		}
		if path[next].Y < y {
			y--
		}
		p := image.Point{x, y}
		ps = append(ps, p)
		if p.Eq(path[next]) {
			next++
		}
		if next >= len(path) {
			return ps
		}
	}
}

func paths(line string) (o []image.Point) {
	nums := regexp.MustCompile(`\d+`).FindAllString(line, -1)
	for i := 0; i <= len(nums)-2; i += 2 {
		o = append(o, image.Point{atoi(nums[i]), atoi(nums[i+1])})
	}
	return o
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
