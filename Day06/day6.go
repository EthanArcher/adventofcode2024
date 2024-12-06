package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {



	mapPlan := readLineSeparatedFile("input.txt")
	var guardRow, guardCol int
	
	// find the guard
	for r := range mapPlan{
		for c := range len(mapPlan[r]) {
			 if mapPlan[r][c] == "^" {
				guardRow = r
				guardCol = c
			 }
		}
	}

	// Part 1 - normal operation
	visited, _ := checkTheRoute(mapPlan[guardRow][guardCol], guardRow, guardCol, mapPlan)
	fmt.Println(len(visited))

	trapCounter := 0

	for _,pos := range visited {

		// Convert to integers and multiply
		parts := strings.Split(pos, ",")
		r, err1 := strconv.Atoi(parts[0])
		c, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			panic("Invalid numbers in mul: " + pos)
		}

		// modifiy the map
		mapPlan[r][c] = "#"
		_, trapped := checkTheRoute("^", guardRow, guardCol, mapPlan)

		if trapped {
			fmt.Println("trapped")
			trapCounter++
		}

		mapPlan[r][c] = "."

	}

	fmt.Println(trapCounter)


}

func checkTheRoute(guardDirection string, guardRow int, guardCol int, mapPlan [][]string) ([]string, bool) {

	visited := []string{fmt.Sprintf("%d,%d", guardRow, guardCol)}
	visitedWithDirection := []string{fmt.Sprintf("%d,%d,%v", guardRow, guardCol, guardDirection)}
	whileLoop := true 
	loop := false
	for whileLoop {
		
		newR, newC := move(guardDirection, guardRow, guardCol)

		if newR == -1 || newR > len(mapPlan)-1 || newC == -1 || newC > len(mapPlan[0])-1 {
			whileLoop = false
			break
		}

		if mapPlan[newR][newC] == "#" {
			guardDirection = rotate(guardDirection)
			continue
		}

		visited, _ = addToSetIfNotThere(visited, fmt.Sprintf("%d,%d", newR, newC))
		visitedWithDirection, loop = addToSetIfNotThere(visitedWithDirection, fmt.Sprintf("%d,%d,%v", newR, newC, guardDirection))

		if loop {
			whileLoop = false
			break
		}

		guardRow, guardCol = newR, newC
	}

	return visited, loop;

	//fmt.Println(visited)

}

func addToSetIfNotThere(s []string, toAdd string) ([]string, bool) {
    for _, v := range s {
        if v == toAdd {
            return s, true // Element already exists, return unchanged
        }
    }
    return append(s, toAdd), false // Add the element and return the updated slice
}

func move(direction string, currentRow int, currentColumn int) (int,int) {

	switch direction {
	case "^":
		return currentRow-1,currentColumn
	case ">":
		return currentRow,currentColumn+1
	case "<":
		return currentRow, currentColumn-1
	case "v":
		return currentRow+1, currentColumn
	default:
		return -1,-1
	}
}

func rotate(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "<":
		return "^"
	case "v":
		return  "<"
	default:
		return ""
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