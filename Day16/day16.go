package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var maxCost int

type position struct {
	r int
	c int
	d direction 
}

type direction struct {
	r int 
	c int 
}

type coordinates struct {
	r int 
	c int 
}

type directionAndCost struct {
	direction position
	cost int
}

var clockwise int = 1
var anticlockwise int = -1

var up direction = direction{-1,0}
var down direction = direction{1,0}
var left direction = direction{0,-1}
var right direction = direction{0,1}

var S coordinates
var E coordinates

func main () {
	grid := readGridFromFile("input.txt")
	S = find("S", grid)
	E = find("E", grid)

	var weights map[position]int = make(map[position]int)
	var reverseWeights map[position]int = make(map[position]int)
	var visited map[position]bool = make(map[position]bool)
	var reverseVisited map[position]bool = make(map[position]bool)

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == "." || grid[r][c] == "S" || grid[r][c] == "E" {
				weights[position{r,c, right}] = math.MaxInt
				weights[position{r,c, left}] = math.MaxInt
				weights[position{r,c, up}] = math.MaxInt
				weights[position{r,c, down}] = math.MaxInt
				visited[position{r,c, right}] = false
				visited[position{r,c, left}] = false
				visited[position{r,c, up}] = false
				visited[position{r,c, down}] = false

				reverseWeights[position{r,c, right}] = math.MaxInt
				reverseWeights[position{r,c, left}] = math.MaxInt
				reverseWeights[position{r,c, up}] = math.MaxInt
				reverseWeights[position{r,c, down}] = math.MaxInt
				reverseVisited[position{r,c, right}] = false
				reverseVisited[position{r,c, left}] = false
				reverseVisited[position{r,c, up}] = false
				reverseVisited[position{r,c, down}] = false
			}
		}
	}

	weights[position{S.r, S.c, right}] = 0

	for !visited[position{E.r, E.c, right}] || !visited[position{E.r, E.c, up}] {
		lowestUnvisitedPosition := findLowestUnvisitedPosition(weights, visited)
		travel(lowestUnvisitedPosition, weights, visited)
	}

	maxCost = int(math.Min(float64(weights[position{E.r, E.c, right}]), float64(weights[position{E.r, E.c, up}])))

	fmt.Println(maxCost)

	// find the reverse paths

	reverseWeights[position{E.r, E.c, down}] = 0
	reverseWeights[position{E.r, E.c, left}] = 0

	for !reverseVisited[position{S.r, S.c, left}] || !reverseVisited[position{S.r, S.c, down}] {
		lowestUnvisitedPosition := findLowestUnvisitedPosition(reverseWeights, reverseVisited)
		travel(lowestUnvisitedPosition, reverseWeights, reverseVisited)
	}

	seats := 0
	for r := range grid {
		for c := range grid[r] {
			minForward := min(weights[position{r, c, right}],weights[position{r, c, up}], weights[position{r, c, left}], weights[position{r, c, down}])
			minReverse := min(reverseWeights[position{r, c, right}],reverseWeights[position{r, c, up}], reverseWeights[position{r, c, left}], reverseWeights[position{r, c, down}])

			// fmt.Println(r,c,minForward,minReverse,minForward+minReverse)

			total := minForward + minReverse
			if int(math.Abs(float64(total))) == maxCost || total + 1000 == maxCost {
				// fmt.Println(r,c,minForward,minReverse,minForward+minReverse)
				seats++
			}
		}
	}
	fmt.Println(seats)

}

func min(a, b, c, d int) int {
    return int(math.Min(math.Min(float64(a), float64(b)), math.Min(float64(c), float64(d))))
}

func findLowestUnvisitedPosition(weights map[position]int, visited map[position]bool) position {
	currentCost := math.MaxInt
	lp := position{}

	for k,v := range visited {
		if !v && weights[k] < currentCost {
			currentCost = weights[k]
			lp = k
		}
	}
	return lp
}

func travel(currentPosition position, weights map[position]int, visited map[position]bool) {

	currentDirection := currentPosition.d
	currentCost := weights[currentPosition]
	visited[currentPosition] = true

	// fmt.Println("visited", currentPosition, "cost", currentCost)

	// move forward one space in the current direction
	np := add(currentPosition, currentDirection)

	cw := position{currentPosition.r, currentPosition.c, rotate(currentDirection, clockwise)}
	acw := position{currentPosition.r, currentPosition.c, rotate(currentDirection, anticlockwise)}

	// Check if the key exists in the map and hasnt been visited yet
	if _, ok := visited[np]; ok && !visited[np] && weights[np] > currentCost + 1 {
		weights[np] = currentCost + 1
	} 

	// move forward 1 space after rotating 90 degrees clockwise
	if _, ok := visited[cw]; ok && !visited[cw] && weights[cw] > currentCost + 1000 {
		weights[cw] = currentCost + 1000
	} 
	
	// move forward 1 space after rotating 90 degrees anticlockwise
	if _, ok := visited[acw]; ok && !visited[acw] && weights[acw] > currentCost + 1000 {
		weights[acw] = currentCost + 1000
	} 

}

func rotate(currentDirection direction, rotation int) direction {
    // Define transitions for each direction and rotation
	// fmt.Println("currentDirection", currentDirection)
    transitions := map[direction]map[int]direction{
        up: {
            anticlockwise: left,
            clockwise:     right,
        },
        right: {
            anticlockwise: up,
            clockwise:     down,
        },
        down: {
            anticlockwise: right,
            clockwise:     left,
        },
        left: {
            anticlockwise: down,
            clockwise:     up,
        },
    }

    // Look up the new direction
    if newDirection, ok := transitions[currentDirection][rotation]; ok {
        return newDirection
    }

    panic("Unknown direction or rotation")
}

func add(a position, b direction) position {
	return position{
		r: a.r + b.r,
		c: a.c + b.c,
		d: b,
	}
}

func find(v string, g [][]string) coordinates{
	for r := range g {
		for c := range g[0]{
			if g[r][c] == v {
				return coordinates{r,c}
			}
		}
	}
	panic("Could not find " + v)
}

func readGridFromFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return lines
}