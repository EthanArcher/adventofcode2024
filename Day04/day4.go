package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var count int = 0
var xmasCount int = 0

func main() {

	input := readLineSeparatedFile("input.txt")

	for r := range input {
		for c := range input[r] {
			if input[r][c] == "X" {
				checkSurrounding(input, r, c)
			} else if input[r][c] == "A" {
				checkForXmas(input, r, c)
			}
		}
	}

	fmt.Println(count)
	fmt.Println(xmasCount)

}

func checkSurrounding(input [][]string, r int, c int) {
	values := []int{-1, 0, 1}
	for _, i := range values{
		if r + (3*i) >= 0 && r + (3*i) < len(input) {
			for _, j := range values {
				if c+ (3*j) >= 0 && c+(3*j) < len(input[r]) {			
					if input[r+i][c+j] == "M" && input[r+(i*2)][c+(j*2)] == "A" && input[r+(i*3)][c+(j*3)] == "S" {
						count++
					}
				}
			}
		}
	}
}

func checkForXmas(input [][]string, r int, c int) {
	if r-1 >= 0 && r+1< len(input) && c-1>=0 && c+1<len(input[r]) {
		if input[r-1][c-1] == "M" && input[r+1][c+1] == "S" && input[r-1][c+1] == "S" && input[r+1][c-1] == "M" {
			xmasCount++
		}
		if input[r-1][c-1] == "M" && input[r+1][c+1] == "S" && input[r-1][c+1] == "M" && input[r+1][c-1] == "S" {
			xmasCount++
		}
		if input[r-1][c-1] == "S" && input[r+1][c+1] == "M" && input[r-1][c+1] == "S" && input[r+1][c-1] == "M" {
			xmasCount++
		}
		if input[r-1][c-1] == "S" && input[r+1][c+1] == "M" && input[r-1][c+1] == "M" && input[r+1][c-1] == "S" {
			xmasCount++
		}
	}
}

// returns a 2d array of [row][column]
func readLineSeparatedFile(filename string) [][]string {
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