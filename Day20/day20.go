package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Item struct {
	position	position
	route 		[]position
	cost 		int 
	index 		int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.index = -1 
	*pq = old[0 : n-1]
	return item
}

type position struct {
	row int
	column int
}

type direction struct {
	row int
	column int
}

var racetrack map[position]bool = make(map[position]bool)
var start position
var end position

var up direction = direction{-1,0}
var down direction = direction{1,0}
var left direction = direction{0,-1}
var right direction = direction{0,1}

var shortestPath int

var directions = []direction{up,right,down,left}

var cheats map[string]bool = make(map[string]bool)

func main() {

	grid := readGridFromFile("input.txt")
	readGrid(grid)
	shortestPaths := findShortestPath(racetrack, []position{start})
	shortestPath = len(shortestPaths[end]) - 1
	fmt.Println("Shortest path without cheating is", shortestPath)
	
	acc := 0
	
	for _,p := range shortestPaths[end] {
		time := 20
		acc += skipsFromPosition(p, time, shortestPaths)
	}
	
	fmt.Println(acc)
}

func skipsFromPosition(p position, time int, shortestPaths map[position][]position) int {

	goodSkips := 0 
	for _,sp := range shortestPaths[end] {
		startAndEnd := fmt.Sprintf("%v, %v -> %v, %v", p.row, p.column,sp.row,sp.column)
		endAndStart := fmt.Sprintf("%v, %v -> %v, %v", sp.row,sp.column,p.row, p.column)

		if cheats[startAndEnd] {
			continue
		}

		cheats[startAndEnd] = true
		cheats[endAndStart] = true

		md := manhattanDistance(p, sp)
		if md <= time {
			diff := int(math.Abs(float64(len(shortestPaths[sp]) - len(shortestPaths[p]))))
			if diff - md >= 100 {
				goodSkips += 1
			}
		}
	}
	return goodSkips
}

func manhattanDistance(p1 position, p2 position) int {
	return int(math.Abs(float64(p1.row) - float64(p2.row)) + math.Abs(float64(p1.column)- float64(p2.column)))
}

func findShortestPath(rt map[position]bool, cheats []position) map[position][]position {
	racetrack := make(map[position]bool)
	for k,v := range rt {
		racetrack[k] = v
	}

	for _,c := range cheats {
		racetrack[c] = true
	}

	pq := make(PriorityQueue, 0) 
	heap.Init(&pq)
	heap.Push(&pq, &Item{
	 position: start,
	 route: []position{start},
	 cost: 0,
	 index: 0,
	})

	visited := make(map[position]int)
	routes := make(map[position][]position)

	for p := range racetrack {
		visited[p] = -1
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.position

		visited[pos] = item.cost
		routes[pos] = item.route

		for _,d := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			newRoute := []position{np}
			newRoute = append(newRoute, item.route...)
			if visited[np] == -1 && racetrack[np]{
				heap.Push(&pq, &Item{
					position: np,
					route: newRoute,
					cost: item.cost + 1,
				})
			}
		}
	}

	return routes

}

func readGrid(grid [][]string) {
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == "S" {
				start = position{row: r, column: c}
				racetrack[position{row: r, column: c}] = true
			} else if grid[r][c] == "E" {
				end = position{row: r, column: c}
				racetrack[position{row: r, column: c}] = true
			} else if grid[r][c] == "." {
				racetrack[position{row: r, column: c}] = true
			}
		}
	}
}

func readGridFromFile(filename string) [][]string{
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open", filename)
		os.Exit(1)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		os.Exit(1)
	}

	return lines

}