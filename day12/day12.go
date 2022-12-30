package day12

import (
	"container/heap"
	"fmt"
	"github.com/Salve/AdventOfCode2022/inputs"
	"github.com/Salve/AdventOfCode2022/registry"
	"math"
	"strings"
)

const day = 12

var input []byte

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1and2()
}

func part1and2() {
	m, start, end := forestmap(string(input))
	distances := dijkstra(m, end)

	part1 := distances[start]
	fmt.Printf("Part 1: %v\n", part1)

	part2 := math.MaxInt
	for p, v := range distances {
		if m[p] == int('a') && v < part2 {
			part2 = v
		}
	}
	fmt.Printf("Part 2: %v\n", part2)
}

func forestmap(input string) (o map[point]int, start, end point) {
	o = make(map[point]int, len(input))
	for y, line := range strings.Split(input, "\n") {
		for x := range line {
			p := point{x, y}
			o[p] = int(line[x])
			if line[x] == 'S' {
				start = p
				o[p] = int('a')
			}
			if line[x] == 'E' {
				end = p
				o[p] = int('z')
			}
		}
	}
	return o, start, end
}

func adjacent(c map[point]int, p point) []point {
	o := make([]point, 0, 4)
	for _, a := range []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}} {
		if _, ok := c[a]; !ok {
			continue
		}
		o = append(o, a)
	}
	return o
}

func dijkstra(c map[point]int, start point) map[point]int {
	n := len(c)
	dist := make(map[point]int, n)
	visited := make(map[point]struct{}, n)
	items := make(map[point]*Item, n)
	q := make(PriorityQueue, n)
	dist[start] = 0
	visited[start] = struct{}{}
	i := 0
	for p := range c {
		if p != start {
			dist[p] = math.MaxInt
		}
		item := &Item{
			p:        p,
			priority: dist[p],
			index:    i,
		}
		items[p] = item
		q[i] = item
		i++
	}
	heap.Init(&q)

	for len(q) > 0 {
		u := heap.Pop(&q).(*Item).p
		visited[u] = struct{}{}
		for _, v := range adjacent(c, u) {
			if _, ok := visited[v]; ok {
				continue
			}
			if c[v] < c[u]-1 { // height difference requirement reversed, since we are traversing from end
				continue
			}
			alt := dist[u] + 1
			if alt < dist[v] && alt > 0 { // could be negative because of maxint and rollover
				dist[v] = alt
				q.update(items[v], v, alt)
			}
		}
	}
	return dist
}

type point struct{ x, y int }

type Item struct {
	p        point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, p point, priority int) {
	item.p = p
	item.priority = priority
	heap.Fix(pq, item.index)
}
