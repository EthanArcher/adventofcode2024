package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	testValue int
	numbers []int
}

func main() {

	lines := readLineSeparatedFile("input.txt")
	equations := []equation{}
	for _,line := range lines {
		l := strings.Split(line, ":")
		v,_ := strconv.Atoi(l[0])
		equations = append(equations, equation{testValue: v, numbers: convertLineToInts(l[1])})
	}

	total := 0
	for _,equation := range equations {
		numbers := combine(equation.numbers)
		if contains(numbers, equation.testValue) {
			total += equation.testValue
		}
	}
	fmt.Println(total)
}

func contains(values []int, value int) bool{
	for _,v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func combine(values []int) []int {
    if len(values) > 1 {
        // Compute the addition and multiplication
        add := append([]int{values[0] + values[1]}, values[2:]...)
        mult := append([]int{values[0] * values[1]}, values[2:]...)

		c,_ := asANumber(fmt.Sprintf("%d%d", values[0], values[1]))
		concat := append([]int{c}, values[2:]...)

        // Combine the results recursively
        result := combine(add)
        result = append(result, combine(mult)...)
		result = append(result, combine(concat)...)
        return result
    }

    // Base case: Return the single value slice
    return []int{values[0]}
}


func asANumber(s string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(s))
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