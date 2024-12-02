package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := readLineSeparatedFile("input.txt")

	safeCount := 0

	for _, line := range input {
		values := convertLineToInts(line)
		isValid := isAValidLine(values)

		if isValid {
			safeCount++
		} else {
			for i := 0; i < len(values); i++ {
				if isAValidLine(newSliceWithValueRemoved(values, i)) {
					safeCount++
					break
				}
			}
		}
	}

	println(safeCount)

}

func newSliceWithValueRemoved(slice []int, s int) []int {
	// Create a copy of the slice to avoid modifying the original
	result := append([]int(nil), slice...) 
	return append(result[:s], result[s+1:]...) 
}

func isAValidLine(values []int) bool {	
	ascending := values[1] > values[0]
	for i := 0; i < len(values)-1; i++ {
		d := math.Abs(float64(values[i+1] - values[i]))
		if (d <= 3) && (d > 0) && ((values[i+1] > values[i]) == ascending) { 
			continue
		} else {
			return false
		}
	}
	return true
}

func convertLineToInts(line string) []int {
	var values []int
	for _, strNum := range strings.Fields(line) { // Use strings.Fields for splitting
		num, err := asANumber(strNum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting '%s' to int: %v\n", strNum, err)
			continue // Skip invalid numbers
		}
		values = append(values, num)
	}
	return values
}

func asANumber(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
}


func readLineSeparatedFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return lines
}