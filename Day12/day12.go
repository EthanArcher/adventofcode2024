package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var visited map[string]bool
var areas map[string]int
var fence map[string]int
var up []int = []int{-1,0}
var down []int = []int{1,0}
var left []int = []int{0,-1}
var right []int = []int{0,1}
var directions [][]int = [][]int{up,right,down,left}
var maxRow int
var maxColumn int
var plot [][]string

func main() {

	visited = make(map[string]bool)
	areas = make(map[string]int)
	fence = make(map[string]int)

	plot = readLineSeparatedFileInto2dArray("input.txt")
	maxRow = len(plot)
	maxColumn = len(plot[0])

	for r, line := range plot {
		for c, _ := range line {
			checkNeighbours([]int{r,c}, fmt.Sprintf("%d:%d", r, c), plot[r][c])
		}
	}

	fmt.Println(areas)
	fmt.Println(fence)

	price := 0 
	for k,v := range areas {
		price += v * fence[k]
	}
	fmt.Println(price)
}

func checkNeighbours(pos []int, zone string, plant string) {

	if visited[fmt.Sprintf("%d:%d", pos[0], pos[1])] {
		return 
	}

	if plot[pos[0]][pos[1]] != plant {
		return
	}

	areas[zone] += 1
	visited[fmt.Sprintf("%d:%d", pos[0], pos[1])] = true

	fmt.Println(pos)
	l := !isValidPosition(pos[0]+left[0], pos[1]+left[1]) || (isValidPosition(pos[0]+left[0], pos[1]+left[1]) && plot[pos[0]+left[0]][pos[1]+left[1]] != plant)
	r := !isValidPosition(pos[0]+right[0], pos[1]+right[1]) || (isValidPosition(pos[0]+right[0], pos[1]+right[1]) && plot[pos[0]+right[0]][pos[1]+right[1]] != plant)
	u := !isValidPosition(pos[0]+up[0], pos[1]+up[1]) || (isValidPosition(pos[0]+up[0], pos[1]+up[1]) && plot[pos[0]+up[0]][pos[1]+up[1]] != plant)
	d := !isValidPosition(pos[0]+down[0], pos[1]+down[1]) || (isValidPosition(pos[0]+down[0], pos[1]+down[1]) && plot[pos[0]+down[0]][pos[1]+down[1]] != plant)

	if l && u {
		fmt.Println("top left corner at", pos)
		fence[zone] += 1
	}
	if l && d {
		fmt.Println("bottom left corner at", pos)
		fence[zone] += 1
	}
	if r && u {
		fmt.Println("top right corner at", pos)
		fence[zone] += 1
	}
	if r && d {
		fmt.Println("bottom right corner at", pos)
		fence[zone] += 1
	}

	// interior corners

	if !l && !u && plot[pos[0]+left[0]+up[0]][pos[1]+left[1]+up[1]] != plant {
		fence[zone] += 1
	}
	if !l && !d && plot[pos[0]+left[0]+down[0]][pos[1]+left[1]+down[1]] != plant {
		fence[zone] += 1
	}
	if !r && !u && plot[pos[0]+right[0]+up[0]][pos[1]+right[1]+up[1]] != plant {
		fence[zone] += 1
	}
	if !r && !d && plot[pos[0]+right[0]+down[0]][pos[1]+right[1]+down[1]] != plant {
		fence[zone] += 1
	}

	for _, d := range directions {
		nr, nc := pos[0]+d[0], pos[1]+d[1]
		if isValidPosition(nr, nc) {
			checkNeighbours([]int{nr,nc}, zone, plant)
		}
	}
}

func isValidPosition(r int, c int) bool {
	return r >=0 && c >= 0 && r < maxRow && c < maxColumn
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