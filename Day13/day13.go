package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`\d+`) 

type equation struct {
	x int
	y int
	ans int
}

func main() {

	equations := readEquationsFromFile("input.txt")
	cost := 0

	for _,eq := range equations {

		inverseMatrix := (eq[0].x * eq[1].y) - (eq[0].y * eq[1].x)

		x := float64(((eq[1].y * eq[0].ans) + (-eq[0].y * eq[1].ans))) / float64(inverseMatrix)
		y := float64(((-eq[1].x * eq[0].ans) + (eq[0].x * eq[1].ans))) / float64(inverseMatrix)

		if  !isWholeNumber(x) || !isWholeNumber(y){
			continue
		}

		cost += int(x)*3
		cost += int(y)
	}

	fmt.Println(cost)

}

func isWholeNumber(num float64) bool {
	return num == math.Floor(num)
}

func readEquationsFromFile(filename string) [][]equation {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var blocks []string
	currentBlock := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // Empty line
			blocks = append(blocks, currentBlock.String())
			currentBlock.Reset() // Start a new block
		} else {
			currentBlock.WriteString(line)
		}
	}
	blocks = append(blocks, currentBlock.String())

	equations := make([][]equation, 0)

	for _,b := range blocks {
		m := re.FindAllString(b, -1)
		eq1 := equation{x:asANumber(m[0]), y:asANumber(m[2]), ans:asANumber(m[4])+10000000000000}
		eq2 := equation{x:asANumber(m[1]), y:asANumber(m[3]), ans:asANumber(m[5])+10000000000000}
		equations = append(equations, []equation{eq1, eq2})
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return equations
}

func asANumber(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Errorf("error converting string '%s' to int: %w", s, err)	
		return 0
	}
	return num
}