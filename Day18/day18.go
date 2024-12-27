package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

var up coordinate = coordinate{0,-1}
var down coordinate = coordinate{0,1}
var left coordinate = coordinate{-1,0}
var right coordinate = coordinate{1,0}

var corruptedCoordinates = make(map[coordinate]bool)
var visited = make(map[coordinate]bool)
var cost = make(map[coordinate]int)
var coordinates []coordinate

var xMax = 70
var yMax = 70

var finished = false
var startingX = 0
var startingY = 0

// An Item is something we manage in a priority queue.
type Item struct {
	coordinate	coordinate	// The value of the item; arbitrary.
	cost 		int    		// The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index 		int 		// The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// return the lowest priority value here
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
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, co coordinate, cost int) {
	item.coordinate = co
	item.cost = cost
	heap.Fix(pq, item.index)
}

func main () {
	coordinates = readCoordinatesFromFile("input.txt")

	for i:=2000; i<len(coordinates); i++ {
		fmt.Println(coordinates[i])
		fmt.Println(i, escape(i))
	}
	
}

func escape(nanoseconds int) bool{

	for c := range cost {
		delete(cost, c)
	}
	for c := range visited {
		delete(visited, c)
	}
	for c := range corruptedCoordinates {
		delete(corruptedCoordinates, c)
	}

	finished := false 

	startingPos := coordinate{x: startingX, y: startingY}
	cost[startingPos] = 0
	
	pq := make(PriorityQueue, 1)
    pq[0] = &Item{
        coordinate: startingPos,
        cost:       0,
        index:      0,
    }
    heap.Init(&pq)
	n := 0
	for n <= nanoseconds { 
		corruptedCoordinates[coordinates[n]] = true
		n++
	}

	for !finished {

		item := heap.Pop(&pq).(*Item)

		if visited[item.coordinate] || !validPosition(item.coordinate) {
			continue
		}

		visited[item.coordinate] = true

		for _,d := range []coordinate{up, down, left, right} {
			newX := item.coordinate.x + d.x
			newY := item.coordinate.y + d.y
			newPos := coordinate{x: newX, y: newY}
			if !visited[newPos] && validPosition(newPos) {
				cost[newPos] = cost[item.coordinate] + 1
				heap.Push(&pq, &Item{coordinate: newPos, cost: cost[newPos]})
			}
			if newPos.x == xMax && newPos.y == yMax {
				finished = true
				fmt.Println("finished cost is", cost[newPos], "for", nanoseconds)
				break
			}
		}
	}

	return finished

}

func validPosition(c coordinate) bool {
	if c.x < 0 || c.y < 0 {
		return false
	}
	if c.x > xMax || c.y > yMax {
		return false
	}

	if corruptedCoordinates[c] {
		return false
	}

	return true
}

func readCoordinatesFromFile(filename string) []coordinate{
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var coordinates []coordinate
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c :=  strings.Split(scanner.Text(), ",")
		coordinates = append(coordinates,coordinate{x: asANumber(c[0]), y: asANumber(c[1])})
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return coordinates

}

func asANumber(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Errorf("error converting string '%s' to int: %w", s, err)	
		return 0
	}
	return num
}
