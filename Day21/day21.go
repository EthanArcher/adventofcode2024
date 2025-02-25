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

var subSequenceCache = make(map[string][]string)
var sequenceLengthCache = make(map[string]int)

var numberOfDpadRobots int = 2

func main() {

	commands := readCommandsFromFile("input-example.txt")
	re := regexp.MustCompile(`\d+`)

	acc := 0

	for _, c := range commands {
		combinedStr := strings.Join(c, "")
		digits := re.FindString(combinedStr)
		value, _ := strconv.Atoi(digits)
		sequenceLength := enterCodeOnNumberPad(c, numberOfDpadRobots)
		fmt.Println(sequenceLength, value)
		acc += sequenceLength * value
	}

	fmt.Println(acc)

}

func enterCodeOnNumberPad(code []string, numberOfRobots int) int {
	previousNumberEntered := "A"
	totalNumberOfCommands := 0

	for _, n := range code {
		numberOfCommands := shortestRouteOnNumberPad(numPad[previousNumberEntered], numPad[n], numberOfRobots)
		previousNumberEntered = n
		totalNumberOfCommands += numberOfCommands
	}

	return totalNumberOfCommands

}

// this method takes a start and end position on the number pad and determines the shortest route
// between the buttons
func shortestRouteOnNumberPad(start position, end position, numberOfRobots int) int {

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		position: start, 
		route: []string{}, 
		cost: 0,
	})

	visited := make(map[position]bool)
	shortestRoute := make(map[position][]string)
	shortestSequence := make(map[position]int)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.position

		if visited[pos] {
			if len(item.route) < len(shortestRoute[pos]) {
				shortestRoute[pos] = item.route
				shortestSequence[pos] = item.cost
			}

			if pos.row == end.row && pos.column == end.column {
				continue
			}

		} else {
			visited[pos] = true
			shortestRoute[pos] = item.route
			shortestSequence[pos] = item.cost
		}

		for d, pad := range directions {
			np := position{row: pos.row + d.row, column: pos.column + d.column}
			if isValidPosition(numPad, np) && !visited[np] {
				newRoute := append(item.route, pad)
				newRouteWithPressA := append(newRoute, "A")
				routeCost := enterSequenceOnDPad(newRouteWithPressA, numberOfRobots)

			   	newItem := Item{
					position: np,
					route: newRoute,
					cost: routeCost,
				}
				heap.Push(&pq, &newItem)
			}
		}
	}
	fmt.Println("shortest route from", start, "to", end, "is", shortestRoute[end], "length", shortestSequence[end])

	return shortestSequence[end]

}

func enterSequenceOnDPad(sequence []string, numberOfRobots int) int {

	sq := strings.Join(sequence, "") + strconv.Itoa(numberOfRobots)

	//Check if the pattern has already been computed
	if result, exists := sequenceLengthCache[sq]; exists {
		return result
	}

	subsequences := [][]string{}
	subsequence := []string{}

	for _,s := range sequence {
		subsequence = append(subsequence, s)
		if s == "A" {
			subsequences = append(subsequences, subsequence)
			subsequence = []string{}
		} 
	}
	subsequences = append(subsequences, subsequence)

	sequenceLength := 0
	for _,ss := range subsequences {
		commandsToEnterSubsequence := enterSubSequenceOnDPad(ss)
		if numberOfRobots == 1 {
			sequenceLength += len(commandsToEnterSubsequence)
		} else {
			sequenceLength += enterSequenceOnDPad(commandsToEnterSubsequence, numberOfRobots - 1)
		}
	}

	sequenceLengthCache[sq] = sequenceLength

	return sequenceLength
}

func enterSubSequenceOnDPad(sequence []string) []string {

	sq := strings.Join(sequence, "")

	// Check if the pattern has already been computed
	if result, exists := subSequenceCache[sq]; exists {
		return result
	}

	press := []string{}

	for i := range len(sequence) {
		if i == 0 {
			press = append(press, enterDirectionOnDPad(dirPad["A"], dirPad[sequence[i]])...)
		} else {
			press = append(press, enterDirectionOnDPad(dirPad[sequence[i-1]], dirPad[sequence[i]])...)
		}
	}

	subSequenceCache[sq] = press
	return press

}

func enterDirectionOnDPad(start position, end position) []string {

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		position: start,
		route:    []string{},
		cost:     0,
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
