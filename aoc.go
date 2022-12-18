package main

import (
	"fmt"
	_ "github.com/Salve/AdventOfCode2022/day1"
	_ "github.com/Salve/AdventOfCode2022/day2"
	_ "github.com/Salve/AdventOfCode2022/day3"
	"github.com/Salve/AdventOfCode2022/registry"
	"time"
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
