package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type antenna struct {
	frequency string
	row int
	column int
}

var maxRows int = 0
var maxCols int = 0

func main() {

	// read the file into a 2d array
	lines := readLineSeparatedFileInto2dArray("input.txt")
	
	// gather all the antennas
	antennas := []antenna{}
	for r,line := range lines {
		for c, v := range line {
			if v!= "." {
				antennas = append(antennas, antenna{frequency: v, row: r, column: c})
			}
		}
	}
	maxRows = len(lines)
	maxCols = len(lines[0])

	// find the resonance positions
	resonance := [][]int{}
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			values := findResonance(antennas[i], antennas[j])
			for _,position := range values {
				resonance = appendIfNotThere(resonance,position)
			}
		}
	}

	fmt.Println(len(resonance))


}

func findResonance(antenna1 antenna, antenna2 antenna) [][]int {

	if antenna1.frequency != antenna2.frequency {
		return nil // Return nil for no resonance
    }

	resonance := [][]int{}

	rowDiff := antenna1.row-antenna2.row
	colDiff := antenna1.column-antenna2.column

	check := true
	multiplier := 0 
	for check {
		rd := rowDiff*multiplier
		cd := colDiff*multiplier

		antinode := []int{antenna1.row + rd, antenna1.column + cd}

		if (isInBounds(antinode)) {
			resonance = append(resonance, antinode)
			multiplier++
		} else {
			check = false
		}
	}

	check = true
	multiplier = 0
	for check {
		rd := rowDiff*multiplier
		cd := colDiff*multiplier

		antinode := []int{antenna2.row - rd, antenna2.column - cd}

		if (isInBounds(antinode)) {
			resonance = append(resonance, antinode)
			multiplier++
		} else {
			check = false
		}
	}

	return resonance

}

// Helper function to check if a point is within the grid bounds
func isInBounds(point []int) bool {
	row, col := point[0], point[1]
	return row >= 0 && row < maxRows && col >= 0 && col < maxCols
}

func appendIfNotThere(values [][]int, z []int) [][]int {
    for _, v := range values {
        if len(v) != 2 || len(z) != 2 || v[0] != z[0] || v[1] != z[1] {
            continue // Skip to the next iteration if any condition is false
        }
        return values // Found a match, return the original slice
    }
    return append(values, z) // No match found, append and return
}

func readLineSeparatedFileInto2dArray(filename string) [][]string {
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