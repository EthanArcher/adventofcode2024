package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var towels = make(map[string]bool)
var seen = make(map[string]int)

func main() {

	patterns := readTowelsFromFile("input.txt")

	for i,p := range patterns {
		fmt.Println("checking pattern", i)
		makeThePattern(p)
	}

	c := 0 
	for _,p := range patterns {	
		// fmt.Println(p, seen[p])
		c = c + seen[p]
	}

	fmt.Println(c)

}

func makeThePattern(pattern string) int {

	// fmt.Println(pattern)

	// Check if the pattern has already been computed
	if result, exists := seen[pattern]; exists {
		// fmt.Println("result seen for", pattern, seen[pattern])
		return result
	}

	// A slice to hold possible remaining patterns that match the prefix
	possibleRemainingPatterns := make([]string, 0, len(towels)) // Preallocate capacity

	n := 0

	// Iterate over all towels to match prefixes and count exact matches
	for t := range towels {
		if pattern == t {
			n++
		}
		if strings.HasPrefix(pattern, t) {
			// Only add the remainder part of the pattern after removing the prefix
			possibleRemainingPatterns = append(possibleRemainingPatterns, pattern[len(t):])
		}
	}

	// fmt.Println("possilbe combos", possibleRemainingPatterns)

	// Recursively compute results for remaining patterns
	for _, p := range possibleRemainingPatterns {
		n += makeThePattern(p)
	}

	// Store the result in the seen map for memoization
	seen[pattern] = n

	return n
}

func readTowelsFromFile(filename string) []string{
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var sections []string
	currentBlock := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // Empty line
			sections = append(sections, currentBlock.String())
			currentBlock.Reset() // Start a new block
		} else {
			currentBlock.WriteString(line+"\n")
		}
	}
	sections = append(sections, currentBlock.String())

	re := regexp.MustCompile(`[a-zA-Z]+`)
	availableTowels := re.FindAllString(sections[0], -1)

	for _,t := range availableTowels {
		towels[t] = true
	}

	patterns := []string{}
	for _, p := range strings.Split(sections[1], "\n") {
		if strings.TrimSpace(p) != "" {
			patterns = append(patterns, strings.TrimSpace(p))
		}
	}

	return patterns

}