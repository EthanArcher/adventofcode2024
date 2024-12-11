package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	column int
}

// Define the directions as constants using iota
const (
    Up = iota
    Down
    Left
    Right
)

// Define a type for directions (optional but recommended)
type Direction int

var numRows int = 0
var numColumns int = 0

func main() {

	mapPlan := readLineSeparatedFileInto2dArray("input.txt")

    startingPositions := make(map[Position][]Position)

	numRows = len(mapPlan)
	numColumns = len(mapPlan[0])

	for r := range mapPlan {
		for c := range len(mapPlan[r]) {
			if mapPlan[r][c] == 0 {
				startingPositions[Position{row: r, column: c}] = []Position{}
			}
		}
	}

	acc := 0
	for sp, _ := range startingPositions {
		peaks := findPeaksThatCanBeReached(mapPlan, sp, []Position{})
		fmt.Println("starting from ", sp, "reached", len(peaks), "peaks")
		acc += len(peaks)
	}


	fmt.Println(acc)

}

func findPeaksThatCanBeReached(mapPlan [][]int, currentPosition Position, peaks []Position) []Position {

	for _, d := range []Direction{Up, Down, Left, Right} {
		np := move(currentPosition, d)

		// Bounds checking
		if np.row < 0 || np.row >= numRows || np.column < 0 || np.column >= numColumns {
			continue
		} 

		if mapPlan[currentPosition.row][currentPosition.column] == 8 && mapPlan[np.row][np.column] == 9 {
			pos := Position{row: np.row, column: np.column}
			// toggle for sets or all
			peaks = appendToSet(peaks, pos)
			//peaks = append(peaks, pos)
		} else if mapPlan[np.row][np.column] == mapPlan[currentPosition.row][currentPosition.column] + 1 {
			peaks = findPeaksThatCanBeReached(mapPlan, np, peaks)
		}

	}
	return peaks

}

func appendToSet(positions []Position, np Position) []Position{

	for _, p := range positions {
		if p == np {
			return positions
		}
	}
	return append(positions, np)
}


// Function to move within a 2D array
func move(p Position, direction Direction) Position {
	
	row := p.row
	col := p.column
	
    switch direction {
	case Up:
        row--
    case Down:
        row++
    case Left:
        col--
    case Right:
        col++
    }
    return Position{row: row, column: col}
}

func asANumber(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
}

func convertLineToInts(line string) []int {
	var values []int
	for _, strNum := range strings.Split(line, "") { // Use strings.Fields for splitting
		num, err := asANumber(strNum)
		if err != nil {
			continue // Skip invalid numbers
		}
		values = append(values, num)
	}
	return values
}

func readLineSeparatedFileInto2dArray(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, convertLineToInts(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return lines
}