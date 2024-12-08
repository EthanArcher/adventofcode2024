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

func main() {

	// read the file into a 2d array
	lines := readLineSeparatedFileInto2dArray("input-example.txt")
	
	// gather all the antennas
	antennas := []antenna{}
	for r,line := range lines {
		for c, v := range line {
			if v!= "." {
				antennas = append(antennas, antenna{frequency: v, row: r, column: c})
			}
		}
		fmt.Println(line)
	}
	fmt.Println(antennas)

	findResonance(antennas[0], antennas[1])

}

func findResonance(antenna1 antenna, antenna2 antenna) []int {

    if antenna1.frequency != antenna2.frequency {
        return []int{} // Return an empty slice if frequencies don't match
    }

	rowDiff := antenna1.row-antenna2.row
	colDiff := antenna1.column-antenna2.column

	fmt.Println(antenna1.row + rowDiff, antenna1.column + colDiff )
	fmt.Println(antenna2.row - rowDiff, antenna2.column - colDiff )

	return []int{}

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