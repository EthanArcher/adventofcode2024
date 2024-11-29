package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(readCommaSeparatedFile("input.txt"))
}

func readCommaSeparatedFile(filename string) []string {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return s
}