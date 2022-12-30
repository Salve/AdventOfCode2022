package main

import (
	"fmt"
	"time"

	_ "github.com/Salve/AdventOfCode2022/day1"
	_ "github.com/Salve/AdventOfCode2022/day10"
	_ "github.com/Salve/AdventOfCode2022/day11"
	_ "github.com/Salve/AdventOfCode2022/day12"
	_ "github.com/Salve/AdventOfCode2022/day13"
	_ "github.com/Salve/AdventOfCode2022/day2"
	_ "github.com/Salve/AdventOfCode2022/day3"
	_ "github.com/Salve/AdventOfCode2022/day4"
	_ "github.com/Salve/AdventOfCode2022/day5"
	_ "github.com/Salve/AdventOfCode2022/day6"
	_ "github.com/Salve/AdventOfCode2022/day7"
	_ "github.com/Salve/AdventOfCode2022/day8"
	_ "github.com/Salve/AdventOfCode2022/day9"
	"github.com/Salve/AdventOfCode2022/registry"
)

func main() {
	name, f := registry.Last()
	fmt.Printf("--- Running last day (%d) ---\n", name)
	timeFunc(f)
}

func timeFunc(f func()) time.Duration {
	start := time.Now()
	f()
	d := time.Now().Sub(start)
	fmt.Printf("--- Execution time: %s ---\n", d)
	return d
}
