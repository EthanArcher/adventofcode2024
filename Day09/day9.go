package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

	values := readLineFromFile("input.txt")

	slice := []int{}

	gaps := 0

	for i, v := range values {
		if i % 2 == 0{
			for j:=0; j<v; j++ {
				slice = append(slice, i/2)
			}
		} else {
			for j:=0; j<v; j++ {
				slice = append(slice, -1)
				gaps++
			}
		}
	}

	part2Slice := append([]int{}, slice...)
	part2SliceBefore := append([]int{}, slice...)
	
	left := 0
	right := len(slice)-1

	for left<right {

		for slice[left] != -1 && left<right {
			left++
		}

		for slice[right] == -1 && right>left {
			right--
		}

		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}

	acc := 0
	for i,v := range slice {
		if slice[i] == -1 {
			break
		}
		acc += i*v
	}
	fmt.Println(acc)

	fmt.Println("Part 2")

	right = len(part2Slice)-1

	for right > 0 {
		if part2SliceBefore[right] == -1 {
			right--
		} else {
			ss := sizeOfSet(right, part2Slice)
			right -= ss
			
			if (right > 0) {
				pos := findSpace(part2Slice[:right+1], ss)
	
				if pos != -1 {
					for x := range ss {
						part2Slice[pos+x], part2Slice[right+x+1] = part2Slice[right+x+1], part2Slice[pos+x]
					}
				}
			}
		}
	}

	acc = 0
	for i,v := range part2Slice {
		if v == -1 {
			continue
		}
		acc += i*v
	}
	
	fmt.Println(acc)
}

func findSpace(slice []int, size int) int {
	r := 0
	for i,v:= range slice {
		if v == -1 {
			r++
		} else {
			r=0
		}
		if r == size {
			return i-(size-1)
		}
	}
	return -1
}

func sizeOfSet(startPosition int, slice []int) int {

	size := 0 
	for startPosition - size >=0 && slice[startPosition] == slice[startPosition - size] {
		size++
	}
	return size

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
	for _, strNum := range strings.Split(line, "") { // Use strings.Fields for splitting
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