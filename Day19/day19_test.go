package main

import (
	"fmt"
	"testing"
)

func TestMakeThePattern(t *testing.T) {
	tests := []struct {
		pattern string
		ways int
	}{
		{pattern: "r", ways: 1},
		{pattern: "rr", ways: 1},
		{pattern: "brr", ways: 2},
		{pattern: "bggr", ways: 1},
		{pattern: "gbbr", ways: 4},
		{pattern: "rrbgbr", ways: 6},
		{pattern: "bwurrg", ways: 1},
	}

	towels = map[string]bool{"r": true, "wr": true, "b": true, "g": true, "bwu": true, "rb": true, "gb": true, "br": true}

	for _, tt := range tests {
		fmt.Println("\n\nStarting test for pattern", tt.pattern)
		
		ways := makeThePattern(tt.pattern)

		if ways != tt.ways {
			t.Errorf("makeThePattern(%s) = %v; want %v", tt.pattern, ways, tt.ways)
		}		
		fmt.Println("End of test for pattern", tt.pattern)

	}
}