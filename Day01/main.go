package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	lines := readLineSeparatedFile("input.txt")
	leftList := []int{}
	rightList := []int{}

	for _, line := range lines {
		left, right, _ := strings.Cut(line, " ")
		leftList = append(leftList, asANumber(left))
		rightList = append(rightList, asANumber(right))
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	acc := 0

	for i := range len(leftList) {
		diff := leftList[i] - rightList[i]
		acc += int(math.Abs(float64(diff))) 	
	}

	fmt.Println(acc)

	sim := 0

	for i := range len(leftList) {
		sim += leftList[i] * countAppearances(leftList[i], rightList)
	}
	
	fmt.Println(sim)

}

func countAppearances(n int, l []int) int {
	count := 0
	for _, num := range l {
		if num == n {
			count++
		}
	}
	return count
}

func asANumber(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Errorf("error converting string '%s' to int: %w", s, err)	
		return 0
	}
	return num
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