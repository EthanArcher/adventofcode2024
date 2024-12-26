package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {

	coordinates := readCoorDinatesFromFile("input-example.txt")
	for _,c := range coordinates {
		fmt.Println(c)
	}

}

func readCoorDinatesFromFile(filename string) []coordinate {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var coordinates []coordinate
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		x,_ := strconv.Atoi(s[0])
		y,_ := strconv.Atoi(s[1])
		coordinates = append(coordinates, coordinate{x: x, y: y})
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	return coordinates
}