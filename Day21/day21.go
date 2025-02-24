package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Item struct {
	position position
	route    []string
	cost     int
	index    int
	driver   []string
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
	row    int
	column int
}

type direction struct {
	row    int
	column int
}

var up direction = direction{-1, 0}
var down direction = direction{1, 0}
var left direction = direction{0, -1}
var right direction = direction{0, 1}

var directions map[direction]string = map[direction]string{
	up:    "^",
	right: ">",
	down:  "v",
	left:  "<",
}

var numPad map[string]position = map[string]position{
	"7": position{row: 0, column: 0},
	"8": position{row: 0, column: 1},
	"9": position{row: 0, column: 2},
	"4": position{row: 1, column: 0},
	"5": position{row: 1, column: 1},
	"6": position{row: 1, column: 2},
	"1": position{row: 2, column: 0},
	"2": position{row: 2, column: 1},
	"3": position{row: 2, column: 2},
	"0": position{row: 3, column: 1},
	"A": position{row: 3, column: 2},
}

var dirPad map[string]position = map[string]position{
	"^": position{row: 0, column: 1},
	"A": position{row: 0, column: 2},
	"<": position{row: 1, column: 0},
	"v": position{row: 1, column: 1},
	">": position{row: 1, column: 2},
}

var dPadCostMap map[string]map[string]int = map[string]map[string]int{
	"A": {
		"^": 1,
		">": 1,
		"v": 2,
		"<": 3,
	},
	"^": {
		"A": 1,
		"v": 1,
		"<": 2,
		">": 2,
	},
	"v": {
		"^": 1,
		"<": 1,
		">": 1,
		"A": 2,
	},
	"<": {
		"v": 1,
		"^": 2,
		">": 2,
		"A": 3,
	},
	">": {
		"v": 1,
		"A": 1,
		"^": 2,
		"<": 2,
	},
}

var dPadCache map[string][]string = make(map[string][]string)

var numberOfDpadRobots int = 2

func main() {

	commands := readCommandsFromFile("input.txt")
	re := regexp.MustCompile(`\d+`)

	acc := 0

	for _, c := range commands {
		combinedStr := strings.Join(c, "")
		digits := re.FindString(combinedStr)
		value, _ := strconv.Atoi(digits)
		commands := enterCodeOnNumberPad(c, numberOfDpadRobots)
		fmt.Println(len(commands), value)
		acc += len(commands) * value
	}

	fmt.Println(acc)

}

func pressingNumberKeypad(sequence []string) []string {
	// fmt.Println("Number pad sequence is", sequence)
	robotCommands := []string{}
	s := "A"
	for _, p := range sequence {
		rm := moveOnNumberPad(s, p)
		// fmt.Println(p, "-->", rm)
		robotCommands = append(robotCommands, rm...)
		s = p
	}
	return robotCommands

}

func moveOnNumberPad(start string, end string) []string {
	sp := numPad[start]
	ep := numPad[end]

	moveToNumber := shortestRoute(sp, ep, numPad)

	return moveToNumber
}

func enterCodeOnNumberPad(code []string, numberOfRobots int) []string {
	humanCommands := []string{}
	previousNumberEntered := "A"

	// fmt.Println("Code is", code)
	for _, n := range code {
		robot1Commands := shortestRouteOnNumberPad(numPad[previousNumberEntered], numPad[n])
		// fmt.Println("Robot 1 commands, previous was", previousNumberEntered, "next is", n, "route is", robot1Commands)
		humanCommands = append(humanCommands, robotsControllingRobots(robot1Commands, numberOfRobots)...)
		previousNumberEntered = n
	}

	return humanCommands
}

func robotsControllingRobots(commands []string, numberOfRobots int) []string{
	sequence := enterSequenceOnDPad(commands)

	if numberOfRobots == 1 {
		return sequence
	} else {
		return robotsControllingRobots(sequence, numberOfRobots - 1)
	}
}

func shortestRouteOnNumberPad(start, end position) []string {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{position: start, route: []string{}, cost: 0})
   
	visited := make(map[position]bool)
	shortestRoute := make(map[position][]string)
   
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
	   	pos := item.position
   
		if visited[pos] { 
			continue // Already processed this node with a shorter path
		}
	   
		visited[pos] = true // Mark as visited AFTER processing neighbors
		shortestRoute[pos] = item.route
	   
		if pos.row == end.row && pos.column == end.column {
		   break // Found the shortest path to the destination
		}
	   
		lastPress := "A"
		if len(item.route) > 0 {
		   lastPress = item.route[len(item.route)-1]
		}
	   
		for d, pad := range directions {
		   	np := position{row: pos.row + d.row, column: pos.column + d.column}
		   	if isValidPosition(numPad, np) && !visited[np] { 
			   	newItem := Item{
					position: np,
					route: append(item.route, pad),
					cost: item.cost + dirPadCost(lastPress, pad),
				}
				heap.Push(&pq, &newItem) 
			}
		}
	}

	// fmt.Println("Shortest route from", start, "to", end, "was", shortestRoute[end])
   
	return append(shortestRoute[end], "A") 
}

// this method takes a start and end position on the number pad and determines the shortest route
// between the buttons
func shortestRouteOnNumberPad2(start position, end position) []string {

	// fmt.Println("Finding shorted route from", start, "to", end)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{position: start, route: []string{}, cost: 0})

	visited := make(map[position]bool)
	shortestRoute := make(map[position][]string)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.position

		lastPress := "A"

		if len(item.route) > 0 {
			lastPress = item.route[len(item.route)-1]
		}

		if visited[pos] {
			if len(item.route) < len(shortestRoute[pos]) {
				shortestRoute[pos] = item.route
			}

			if pos.row == end.row && pos.column == end.column {
				continue
			}

		} else {
			visited[pos] = true
			shortestRoute[pos] = item.route
		}

		for d, pad := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			if isValidPosition(numPad, np) && !visited[np] {
				newItem := Item{
					position: np,
					route:    append(item.route, pad),
					cost:     item.cost + dirPadCost(lastPress, pad),
				}
				heap.Push(&pq, &newItem)
			}
		}
	}


	return append(shortestRoute[end], "A")

}

func enterSequenceOnDPad(sequence []string) []string {
	subsequences := [][]string{}
	subsequence := []string{}

	for _,s := range sequence {
		subsequence = append(subsequence, s)
		if s == "A" {
			subsequences = append(subsequences, subsequence)
			subsequence = []string{}
		} 
	}

	commands := []string{}
	for _,ss := range subsequences {
		commands = append(commands, enterSubSequenceOnDPad(ss)...)
	}

	// fmt.Println(sequence, "is", commands)

	return commands

}

func enterSubSequenceOnDPad(sequence []string) []string {

	press := []string{}

	for i := range len(sequence) {
		if i == 0 {
			press = append(press, enterDirectionOnDPad(dirPad["A"], dirPad[sequence[i]])...)
		} else {
			press = append(press, enterDirectionOnDPad(dirPad[sequence[i-1]], dirPad[sequence[i]])...)
		}
	}

	return press

}

func enterDirectionOnDPad(start position, end position) []string {

	// fmt.Println("Finding shorted route from", start, "to", end)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		position: start,
		route:    []string{},
		cost:     0,
		index:    0,
	})

	visited := make(map[position]bool)
	shortestRoute := make(map[position][]string)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.position

		lastPress := "A"

		if len(item.route) > 0 {
			lastPress = item.route[len(item.route)-1]
		}

		if visited[pos] {
			if len(item.route) < len(shortestRoute[pos]) {
				shortestRoute[pos] = item.route
			}

			if pos.row == end.row && pos.column == end.column {
				continue
			}

		} else {
			visited[pos] = true
			shortestRoute[pos] = item.route
		}

		for d, pad := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			if isValidPosition(dirPad, np) && !visited[np] {

				newItem := Item{
					position: np,
					route:    append(item.route, pad),
					cost:     item.cost + dirPadCost(lastPress, pad),
				}

				heap.Push(&pq, &newItem)
			}
		}
	}

	return append(shortestRoute[end], "A")
}

// Shortest route on the number pad
func shortestRoute(start position, end position, positions map[string]position) []string {

	// fmt.Println("Finding shorted route from", start, "to", end)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		position: start,
		route:    []string{},
		cost:     0,
		index:    0,
		driver:   []string{},
	})

	visited := make(map[position]bool)
	shortestRoute := make(map[position][]string)
	shortestDriverRoute := make(map[position][]string)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.position

		lastPress := "A"

		if len(item.route) > 0 {
			lastPress = item.route[len(item.route)-1]
		}

		if visited[pos] {
			if len(item.driver) < len(shortestDriverRoute[pos]) {
				shortestRoute[pos] = item.route
				shortestDriverRoute[pos] = item.driver
			}

			if pos.row == end.row && pos.column == end.column {
				continue
			}

		} else {
			visited[pos] = true
			shortestRoute[pos] = item.route
			shortestDriverRoute[pos] = item.driver
		}

		for d, pad := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			if isValidPosition(positions, np) && !visited[np] {

				// fmt.Println("\tRoute so far is", item.route, "current position is", pos, "next direction is", pad, "to position", np)

				sdpm := shortestDirectionPadMove(dirPad[lastPress], dirPad[pad], numberOfDpadRobots)

				newItem := Item{
					position: np,
					route:    append(item.route, pad),
					cost:     item.cost + len(sdpm),
					driver:   append(item.driver, sdpm...),
				}

				heap.Push(&pq, &newItem)
			}
		}
	}

	shortest := shortestRoute[end]
	lastPress := shortest[len(shortest)-1]

	// fmt.Println("Shortest route from", start, "to", end, "was", shortestRoute[end])

	// fmt.Println("now finding shortest route from", dirPad[lastPress], "to A")

	pressA := shortestDirectionPadMove(dirPad[lastPress], dirPad["A"], numberOfDpadRobots)
	return append(shortestDriverRoute[end], pressA...)
}

// Shortest route on the dpad
func shortestDirectionPadMove(start position, end position, level int) []string {

	// fmt.Println(level)

	startEndString := fmt.Sprintf("(%d, %d, %d, %d, %d)", start.row, start.column, end.row, end.column, level)

	if _, ok := dPadCache[startEndString]; ok {
		return dPadCache[startEndString]
	}

	if start == end {
		return []string{"A"}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		position: start,
		route:    []string{},
		cost:     0,
		index:    0,
		driver:   []string{},
	})

	visited := make(map[position]bool)
	routes := make(map[position][]string)

	finished := false

	for !finished {
		item := heap.Pop(&pq).(*Item)
		pos := item.position
		visited[pos] = true
		routes[pos] = item.route

		if pos.row == end.row && pos.column == end.column {
			finished = true

			route := []string{}
			if level > 1 {
				robot3FinishPosition := item.route[len(item.route)-1]
				pressA := shortestDirectionPadMove(dirPad[robot3FinishPosition], dirPad["A"], level-1)
				route = append(item.driver, pressA...)

			} else {
				route = append(routes[end], "A")
			}
			dPadCache[startEndString] = route
			return route

		}

		for d, pad := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			if isValidPosition(dirPad, np) && !visited[np] {

				robot3Position := dirPad["A"]
				if len(item.route) > 0 {
					robot3Position = dirPad[item.route[len(item.route)-1]]
				}

				robotMoves := []string{}

				if level > 1 {
					robotMoves = shortestDirectionPadMove(robot3Position, dirPad[pad], level-1)
				} else {
					robotMoves = []string{pad}
				}

				newRoute := append([]string(nil), append(item.route, pad)...)
				newCost := item.cost + len(robotMoves)

				heap.Push(&pq, &Item{
					position: np,
					route:    newRoute,
					cost:     newCost,
					driver:   append(item.driver, robotMoves...),
				})
			}
		}
	}

	return []string{}

}

func isValidPosition(m map[string]position, target position) bool {
	for _, pos := range m {
		if pos.row == target.row && pos.column == target.column {
			return true
		}
	}
	return false
}

func readCommandsFromFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not open", filename)
		os.Exit(1)
	}
	defer file.Close()

	var commands [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		os.Exit(1)
	}

	return commands
}

func dirPadCost(start string, end string) int {
	if start == end {
		return 0
	} else if costs, ok := dPadCostMap[start]; ok {
		if cost, ok := costs[end]; ok {
			return cost
		}
	}
	return -1
}
