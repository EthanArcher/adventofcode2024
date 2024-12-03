package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)

func main() {

	line := readLineFromFile("input.txt")
	matches := pattern.FindAllString(line, -1)

	sum := 0
	doing := true

	for _, m := range matches {
		switch m {
		case "do()":
			doing = true
		case "don't()":
			doing = false
		default:
			if doing {
				sum += multiplyValues(m)
			}
		}
	}

	fmt.Println(sum)

}

// Parses and multiplies values from a "mul(x,y)" string
func multiplyValues(expr string) int {
	// Remove "mul(" and ")" and split by comma
	values := strings.TrimPrefix(strings.TrimSuffix(expr, ")"), "mul(")
	parts := strings.Split(values, ",")
	if len(parts) != 2 {
		panic("Invalid mul format: " + expr)
	}

	// Convert to integers and multiply
	x, err1 := strconv.Atoi(parts[0])
	y, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		panic("Invalid numbers in mul: " + expr)
	}

	return x * y
}

// Reads the first line from a file
func readLineFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines[0]
}
