package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

var up position = position{-1,0}
var down position = position{1,0}
var left position = position{0,-1}
var right position = position{0,1}

func main() {

	warehouse, directions := readFromFile("input.txt")
	fmt.Println("Starting warehouse - Part 1")
	printWarehouse(warehouse)

	rows := len(warehouse)
	cols := len(warehouse[0])
	
	biggerWarehouse := make([][]string, rows)
	for i := range biggerWarehouse {
		biggerWarehouse[i] = make([]string, cols*2)
	}

	for r:= 0; r<len(warehouse); r++ {
		for c:= 0; c<len(warehouse[0]); c++ {
			nc := c*2
			if warehouse[r][c] == "O" {
				biggerWarehouse[r][nc] = "["
				biggerWarehouse[r][nc+1] = "]"
			} else if warehouse[r][c] == "#" {
				biggerWarehouse[r][nc] = "#"
				biggerWarehouse[r][nc+1] = "#"
			} else if warehouse[r][c] == "." {
				biggerWarehouse[r][nc] = "."
				biggerWarehouse[r][nc+1] = "."
			} else if warehouse[r][c] == "@" {
				biggerWarehouse[r][nc] = "@"
				biggerWarehouse[r][nc+1] = "."
			}
		}
	}

	robotPosition := findStartingPosition(warehouse)

	for _,d := range directions {
		move(&robotPosition, getDirection(d), warehouse)
	}

	fmt.Println("\nFinished warehouse")
	printWarehouse(warehouse)

	sum := 0
	for r := range warehouse {
		for c := range warehouse[r] {
			if warehouse[r][c] == "O" {
				sum += (100*r) + c
			}
		}
	}
	fmt.Println(sum)

	fmt.Println("\nBigger warehouse - Part 2")
	
	robotPosition = findStartingPosition(biggerWarehouse)

	for _,d := range directions {
		moveBigBox(&robotPosition, getDirection(d), biggerWarehouse)
	}
	printWarehouse(biggerWarehouse)

	sum = 0
	for r := range biggerWarehouse {
		for c := range biggerWarehouse[r] {
			if biggerWarehouse[r][c] == "[" {
				sum += (100*r) + c
			}
		}
	}
	fmt.Println(sum)

}

func move(currentPosition *position, direction position, warehouse [][]string) {
	nextPosition := nextEmptyPosition(*currentPosition, direction, warehouse)
	if nextPosition.x != -1 && nextPosition.y != -1 {
		warehouse[nextPosition.x][nextPosition.y] = "O"
		warehouse[currentPosition.x][currentPosition.y] = "."
		currentPosition.x = currentPosition.x + direction.x
		currentPosition.y = currentPosition.y + direction.y
		warehouse[currentPosition.x][currentPosition.y] = "@"
	}
}


func moveBigBox(currentPosition *position, direction position, warehouse [][]string) {

	if warehouse[currentPosition.x+direction.x][currentPosition.y+direction.y] == "." {
		// this is just an empty space we can move into
		moveRobot(currentPosition, direction, warehouse)
	} else if warehouse[currentPosition.x+direction.x][currentPosition.y+direction.y] == "#" {
		// do nothing this is a wall
	} else if direction == right || direction == left {
		// move the boxes left or right
		nextPosition := nextEmptyPosition(*currentPosition, direction, warehouse)

		if nextPosition.x != -1 && nextPosition.y != -1 {
			y := nextPosition.y
			for y != currentPosition.y {
				warehouse[currentPosition.x][y] = warehouse[currentPosition.x][y-direction.y]
				y-=direction.y
			}
			moveRobot(currentPosition, direction, warehouse)
		}
	} else if direction == up || direction == down {

			boxPositions := boxPositions(*currentPosition, direction, warehouse)

			if canBumpBoxes(boxPositions, direction, warehouse) {
				bumpBoxes(boxPositions, direction, warehouse)
				moveRobot(currentPosition, direction, warehouse)
			}
	}
}

func moveRobot(currentPosition *position, direction position, warehouse [][]string) {
	warehouse[currentPosition.x+direction.x][currentPosition.y+direction.y] = "@"
	warehouse[currentPosition.x][currentPosition.y] = "."
	currentPosition.x = currentPosition.x + direction.x
	currentPosition.y = currentPosition.y + direction.y
}

func canBumpBoxes(positions []position, direction position, warehouse [][]string) bool {
	canBump := true
	for _,p := range positions {
		newPos := position{p.x+direction.x,p.y+direction.y}
		np := warehouse[newPos.x][newPos.y]

		if np == "[" || np == "]" {
			canBump = canBump && canBumpBoxes(boxPositions(p, direction, warehouse), direction, warehouse)
		} else if np == "." {
			canBump = canBump && true
		} else if np == "#" {
			canBump = canBump && false
		}
	}
	return canBump
}


func getDirection(d string) position {
	switch d {
	case ">":
		return right
	case "^":
		return up
	case "v":
		return down
	case "<":
		return left
	}
	panic("not a valid direction")
}

func nextEmptyPosition(currentPosition position, direction position, warehouse [][]string) position{
	nextPosition := position{currentPosition.x + direction.x, currentPosition.y + direction.y}
	var np position
	if warehouse[nextPosition.x][nextPosition.y] == "." {
		np = nextPosition
	} else if warehouse[nextPosition.x][nextPosition.y] == "#"{
		np = position{-1,-1}
	} else {
		np = nextEmptyPosition(nextPosition, direction, warehouse)
	}
	return np
}

func bumpBoxes(boxes []position, direction position, warehouse [][]string) {
	for _,p := range boxes {
		newPos := position{p.x+direction.x,p.y+direction.y}
		np := warehouse[newPos.x][newPos.y]
		if np == "[" || np == "]" {
			touching := boxPositions(p, direction, warehouse)
			bumpBoxes(touching, direction, warehouse)
		} 
		warehouse[newPos.x][newPos.y] = warehouse[p.x][p.y]
		warehouse[p.x][p.y] = "."		
	}
}

func boxPositions(currentPosition position, direction position, warehouse [][]string) []position {
	if warehouse[currentPosition.x+direction.x][currentPosition.y+direction.y] == "[" {
		return []position{position{currentPosition.x+direction.x, currentPosition.y+direction.y},position{currentPosition.x+direction.x, currentPosition.y+direction.y+1}}
	} else if warehouse[currentPosition.x+direction.x][currentPosition.y+direction.y] == "]" {
		return []position{position{currentPosition.x+direction.x, currentPosition.y+direction.y},position{currentPosition.x+direction.x, currentPosition.y+direction.y-1}}
	} else {
		return []position{}
	}
}

func printWarehouse(warehouse [][]string) {
	for _,l := range warehouse {
		fmt.Println(l)
	}
}

func findStartingPosition(warehouse [][]string) position{
	for r := range warehouse {
		for c := range warehouse[r] {
			if warehouse[r][c] == "@" {
				return position{r,c}
			}
		}
	}
	panic("couldn't find the robot")
}

func readFromFile(filename string) ([][]string, []string){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		panic("Invalid file name")
	}
	defer file.Close()

	var blocks [][]string
	currentBlock := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // Empty line
			blocks = append(blocks, currentBlock)
			currentBlock = []string{}
		} else {
			currentBlock = append(currentBlock, line)
		}
	}
	blocks = append(blocks, currentBlock)

	directions := strings.Builder{}
	for _,l := range blocks[1] {
		directions.WriteString(l)
	}

	warehouse := [][]string{}
	for _,l := range blocks[0] {
		warehouse = append(warehouse, strings.Split(l, ""))
	}

	return warehouse, strings.Split(directions.String(), "")
}