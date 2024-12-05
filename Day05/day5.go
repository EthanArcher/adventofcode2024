package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rules := []string{}
	updates := [][]int{}

	lines := readInputFile("input.txt")

	for _, line := range lines {
		if strings.Contains(line,"|") {
			rules = append(rules, line)
		} else if strings.Contains(line,",") {
			updates = append(updates, splitAndConvert(line, ","))
		}
	}

	counter := 0
	invalidLines := [][]int{}

	for _,update := range updates {

		lineIsValid,_ := isLineValid(update, rules)

		if lineIsValid {
			counter += update[(len(update)-1) / 2]
		} else {
			invalidLines = append(invalidLines, update)
		}
	}

	fmt.Println(counter)

	fixedCounter := 0

	for _,line := range invalidLines {
		lineIsValid,pos := isLineValid(line, rules)
		for !lineIsValid {
			swapValuesAtPosition(line, pos)
			lineIsValid,pos = isLineValid(line, rules)
		}
		fixedCounter += line[(len(line)-1) / 2]
	}

	fmt.Println(fixedCounter)

}

func swapValuesAtPosition(line []int, pos int) {
	line[pos], line[pos+1] = line[pos+1], line[pos]
}

// check is the line valid
func isLineValid(line []int, rules []string) (bool,int) {
	for j := range len(line)-1 {			
		toBeTrue := fmt.Sprintf("%d|%d", line[j], line[j+1])
		toBeFalse := fmt.Sprintf("%d|%d", line[j+1], line[j])

		if contains(rules, toBeTrue) && !contains(rules, toBeFalse){

		} else {
			return false,j
		}
	}
	return true,0
}

// check does a slice contain a value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// spit the line using the separator and return the converted into ints
func splitAndConvert(line string, separator string) []int {
	splitArray := strings.Split(line, separator)
	conv := []int{}
	for _,v := range splitArray {
		c, err := strconv.Atoi(v)
		if err != nil {
			panic("Invalid numbers:" + line)
		}
		conv = append(conv, c)
	}
	return conv
}

// returns an array of lines
func readInputFile(filename string) []string {
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