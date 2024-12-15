package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "time"
)

type robot struct {
	startingX int
	startingY int
	xDir      int
	yDir      int
}

// var xMax int = 11
// var yMax int = 7
var xMax int = 101
var yMax int = 103

func main() {

	positions := make(map[[2]int]int)
	robots := readRobotsFromFile("input.txt")
	for _,robot := range robots {
		xPos, yPos := positionAfterSeconds(robot, 100)
		positions[[2]int{xPos, yPos}] += 1
	}

	xMid := float64(xMax - 1) / 2
	yMid := float64(yMax - 1) / 2
		
	quadrants := make(map[string]int)
	for p,c := range positions {
		switch {
		case float64(p[0]) < xMid && float64(p[1]) < yMid:
			quadrants["TL"] += c
		case float64(p[0]) > xMid && float64(p[1]) < yMid:
			quadrants["TR"] += c
		case float64(p[0]) < xMid && float64(p[1]) > yMid:
			quadrants["BL"] += c
		case float64(p[0]) > xMid && float64(p[1]) > yMid:
			quadrants["BR"] += c
		}
	}

	safetyFactor := 1
	for _,v := range quadrants {
		safetyFactor = safetyFactor * v
	}
	fmt.Println(safetyFactor)

	for i:= 0; i<10000; i++ {
		// fmt.Println()
		// fmt.Println(i, "seconds")
		// fmt.Println()
		for k := range positions {
			delete(positions, k)
		}
			
		for _,robot := range robots {
			xPos, yPos := positionAfterSeconds(robot, i)
			positions[[2]int{xPos, yPos}] += 1
		}
		printRobots(positions, i)
	}
}

func printRobots(positions map[[2]int]int, time int) {
	display := make([][]string, yMax) // Fixed size array
	maybeTree := false
	for y := 0; y < yMax; y++ {
		display[y] = make([]string, xMax)
    	for x := 0; x < xMax; x++ {
			display[y][x] = "."
		}
    }
	for p,_ := range positions {
		display[p[1]][p[0]] = "#"
	}
	for _,r := range display {
		count := 0
		for _,q := range r {
			if q == "#" {
				count++
			}
		}
		if count >30 {
			maybeTree = true
		}
		count = 0
	}

	if maybeTree {
		fmt.Println("\n Seconds", time)
		for _,r := range display {
			fmt.Println(r,)
		}
	}
	//time.Sleep(100 * time.Millisecond) 
}

func positionAfterSeconds(robot robot, seconds int) (int,int) {
	xPos := ((robot.startingX + (robot.xDir * seconds))%xMax + xMax) % xMax
	yPos := ((robot.startingY + (robot.yDir * seconds))%yMax + yMax) % yMax
	return xPos, yPos
}

func readRobotsFromFile(filename string) []robot {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	re := regexp.MustCompile(`-?\d+`)
	var robots []robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		var numbers []int
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
		robots = append(robots, robot{startingX: numbers[0], startingY: numbers[1], xDir: numbers[2], yDir: numbers[3]})
	}
	return robots
}