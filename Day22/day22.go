package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var step1Cache = make(map[int]int)
var step2Cache = make(map[int]int)
var step3Cache = make(map[int]int)

var potentialSequences = make(map[string]bool)
var sequencesAndPrices = []map[string]int{}


func main() {
	initialBuyerSecrets, _ := readFromFile("input.txt")
	acc := 0
	for _, v := range initialBuyerSecrets {
		acc += nthSecretNumber(v, 2000)
		sequencesAndPrice := findSequenceOfChanges(v)
		sequencesAndPrices = append(sequencesAndPrices, sequencesAndPrice)
	}
	fmt.Println(acc)

	largestTotal := 0
	for ps, _ := range potentialSequences {
		acc := 0
		for i := range initialBuyerSecrets {
			acc += sequencesAndPrices[i][ps]
		}
		if acc > largestTotal {
			largestTotal = acc
		}
	}
	fmt.Println(largestTotal)
}

func findSequenceOfChanges(secret int) map[string]int {
	changes := make(map[string]int)
	seq := []int{}
	current := secret
	previous := 0
	diff := 0

	for range 4 {
		previous = current
		current = nextSecretNumber(current)
		diff = (current % 10) - (previous % 10)
		seq = append(seq, diff)
	}

	for i := 4; i < 2000; i++ {
		addPotentialSequence(seq)
		if _, ok := changes[fmt.Sprintf("%v", seq)]; ok {
			// do nothing 
		} else {
			changes[fmt.Sprintf("%v", seq)] = (current % 10)
		}

		previous = current
		current = nextSecretNumber(current)
		diff = (current % 10) - (previous % 10)

		// remove first element from seq
		seq = seq[1:]
		seq = append(seq, diff)
	}

	return changes

}

func addPotentialSequence(seq []int) {
	// if potentialSequences contains sequence
	if _, ok := potentialSequences[fmt.Sprintf("%v", seq)]; ok {
		return
	} else {
		potentialSequences[fmt.Sprintf("%v", seq)] = true
	}
}


func nthSecretNumber(secret int, times int) int {
	s := secret
	for i := 0; i < times; i++ {
		s = nextSecretNumber(s)
	}
	return s
}

func nextSecretNumber(secret int) int {
	step1 := step1(secret)
	step2 := step2(step1)
	step3 := step3(step2)
	return step3
}

func step1(secret int) int {

	//Check if the pattern has already been computed
	if result, exists := step1Cache[secret]; exists {
		return result
	}

	r := secret * 64
	m := mix(r, secret)
	p := prune(m)

	step1Cache[secret] = p
	return p
}

func step2(secret int) int {

	//Check if the pattern has already been computed
	if result, exists := step2Cache[secret]; exists {
		return result
	}

	r := secret / 32
	m := mix(r, secret)
	p := prune(m)

	step2Cache[secret] = p
	return p
}

func step3(secret int) int {

	//Check if the pattern has already been computed
	if result, exists := step3Cache[secret]; exists {
		return result
	}

	r := secret * 2048
	m := mix(r, secret)
	p := prune(m)

	step3Cache[secret] = p
	return p
}


func mix(input, secret int) int {
	return input ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}


func readFromFile(filename string) ([]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()

    var numbers []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
        if err != nil {
            return nil, fmt.Errorf("error converting line to number: %w", err)
        }
        numbers = append(numbers, num)
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading file: %w", err)
    }

    return numbers, nil
}