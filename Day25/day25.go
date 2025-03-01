package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	keys, locks := readFile("input.txt")

	// fmt.Println("keys", keys)
	// fmt.Println("locks", locks)

	acc :=0

	for _, key := range keys {
		for _, lock := range locks {
			if keyFitsLock(key, lock) {
				fmt.Println("Key fits lock")
				acc++
			}
		}
	}
	fmt.Println(acc)

}

func keyFitsLock(key []int, lock []int) bool {
	for i := 0; i < 5; i++ {
		if key[i] + lock[i] > 5{
			return false
		}
	}
	return true
}

func readFile(filename string) ([][]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		os.Exit(1)
	}
	defer file.Close()

	keys := [][]int{}
	locks := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := []string{}
		for i := 0; i < 7; i++ {
			entry = append(entry, scanner.Text())
			scanner.Scan()
		}
		sums := []int{0,0,0,0,0}
		for i := 1; i < 6; i++ {
			split := strings.Split(entry[i], "")
			for j:=0; j<5;j++ {
				if split[j] == "#" {
					sums[j] += 1
				}
			}
		}
		if entry[0] == "#####" {	
			locks = append(locks, sums)
		} else {
			keys = append(keys, sums)
		}
		scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		os.Exit(1)
	}

	return keys, locks
}