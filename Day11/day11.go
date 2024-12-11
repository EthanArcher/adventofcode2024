package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var resultMap map[int][]int
var stoneCounter map[int]int

func main() {

	resultMap = make(map[int][]int)
	stoneCounter = make(map[int]int)
	stones := readLineFromFile("input.txt")

	for _, s := range stones {
		addToCounter(s, 1)
	}

    for i := 0; i < 75; i++ {
		og := copyMap(stoneCounter)
        for stoneValue, count := range og {
			if count > 0 {
				for _,s := range findStonesCreated(stoneValue) {
					addToCounter(s, count)
				}
				stoneCounter[stoneValue] -= count
			}
        }
    }


	acc := 0
	for _, count := range stoneCounter {
		acc += count
	}

	fmt.Println(acc)

}

func copyMap(originalMap map[int]int) map[int]int {
	newMap := make(map[int]int)
    for key, value := range originalMap {
        newMap[key] = value
    }
    return newMap
}

func addToCounter(stoneValue int, count int) {
	if _, ok := stoneCounter[stoneValue]; !ok {
		stoneCounter[stoneValue] = count
	} else {
		stoneCounter[stoneValue] += count
	}
}


func findStonesCreated(stoneValue int) []int{
    if _, ok := resultMap[stoneValue]; !ok {
        resultMap[stoneValue] = blinkAtStone(stoneValue)
    } 
	return resultMap[stoneValue]
}
	
func blinkAtStone(stoneValue int) []int {
	
	if stoneValue == 0 {
		return []int{1}
	}

	nod := numberOfDigits(stoneValue)
	if nod % 2 == 0 {
		return splitValueInHalf(stoneValue, nod/2)
	}

	return []int{stoneValue * 2024}

}

func splitValueInHalf(value int, p int) []int {
    left := value / int(math.Pow10(len(strconv.Itoa(value))-p))
    right := value % int(math.Pow10(len(strconv.Itoa(value))-p))
    return []int{left, right}
}

func numberOfDigits(value int) int {
	str := strconv.Itoa(value)
	return len(strings.Split(str, ""))
}


// Reads the first line from a file
func readLineFromFile(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	values := convertLineToInts(string(data))
	return values
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