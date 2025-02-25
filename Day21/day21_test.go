package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// func TestMoveOnNumPad_robot1(t *testing.T) {
// 	tests := []struct {
// 		start string
// 		end   string
// 		move  []string
// 	}{
// 		// {start: "7", end: "9", move: []string{">", ">", "A"}},
// 		// {start: "7", end: "5", move: []string{"v", ">", "A"}},
// 		// {start: "7", end: "A", move: []string{"v", "v", "v", ">", ">", "A"}},
// 		// {start: "A", end: "7", move: []string{"^", "^", "^", "<", "<", "A"}},
// 		// {start: "A", end: "0", move: []string{"<", "A"}},
// 		// {start: "A", end: "3", move: []string{"^", "A"}},
// 		{start: "3", end: "7", move: []string{"<", "<", "^", "^", "A"}},
// 		// {start: "7", end: "9", move: []string{">", ">", "A"}},
// 		// {start: "9", end: "A", move: []string{"v", "v", "v", "A"}},
// 	}

// 	for _, tt := range tests {

// 		commands := shortestRouteOnNumberPad(numPad[tt.start], numPad[tt.end], 2)
// 		fmt.Println("Moving from", tt.start, "to", tt.end, "is", commands)

// 		if commands != len(tt.move) {
// 			t.Errorf("moveOnNumberPad(%v, %v) = %v; want %v", tt.start, tt.end, commands, tt.move)
// 		}
// 	}

// 	for _, tt := range tests {
// 		move := shortestRouteOnNumberPad(numPad[tt.start], numPad[tt.end])
// 		fmt.Println("Moving from", tt.start, "to", tt.end, "is", move)
// 		fmt.Println("Expected:", tt.move)
// 		fmt.Println("Actual:  ", move)
	
// 		if len(move) != len(tt.move) || !equalSlices(move, tt.move) {
// 			t.Errorf("moveOnNumberPad(%v, %v) = %v; want %v", tt.start, tt.end, move, tt.move)
// 		}
// 	}
// }

func TestValidPosition(t *testing.T) {
	tests := []struct {
		pos position
		expected bool
	}{
		{pos: position{row: 0, column: 0}},
	}

	for _, tt := range tests {

		v := isValidPosition(dirPad, tt.pos)

		if v != tt.expected {
			t.Errorf("isValidPosition(%v) = %v; want %v", tt.pos, v, tt.expected)
		}
	}
}

func TestEnterDirectionsOnDPad(t *testing.T) {
	tests := []struct {
		start string
		end   string
		move  []string
	}{
		{start: "A", end: "<", move: []string{"v", "<", "<", "A"}},
		{start: "<", end: "A", move: []string{">", ">", "^", "A"}},
	}

	for _, tt := range tests {

		move := enterDirectionOnDPad(dirPad[tt.start], dirPad[tt.end])
		fmt.Println("Moving from", tt.start, "to", tt.end, "is", move)

		if len(move) != len(tt.move) {
			t.Errorf("moveOnNumberPad(%v, %v) = %v; want %v", tt.start, tt.end, move, tt.move)
		}
	}
}

func TestEnterSubSequenceOnDPad(t *testing.T) {
	tests := []struct {
		subSequence []string
		move     []string
	}{
		{subSequence: []string{"<", "A"}, move: []string{"v", "<", "<", "A", ">", ">", "^", "A"}},
		{subSequence: []string{"^", "A"}, move: []string{"<", "A", ">", "A"}},
		{subSequence: []string{"<", "<", "^", "^", "A"}, move: []string{"v", "<", "<", "A", "A", ">", "^", "A", "A", ">", "A"}},
		{subSequence: []string{"^", "^", "<", "<", "A"}, move: []string{"<", "A", "A", "v", "<", "A", "A", ">", ">", "^", "A"}},
		{subSequence: []string{">", ">", "A"}, move: []string{"v", "A", "A", "^", "A"}},
		{subSequence: []string{"v", "v", "v", "A"}, move: []string{"v", "<", "A", "A", "A", "^", ">", "A"}},
		{subSequence: []string{"v", "<", "<", "A", ">", ">", "^", "A"}, move: []string{"v", "<", "A", "<", "A", "A", ">", ">", "^", "A", "v", "A", "A", "^", "<", "A", ">", "A"}},
		{subSequence: []string{"v", "<", "<", "A", "A", ">", "^", "A", "A", ">", "A"}, move: []string{"v", "<", "A", "<", "A", "A", ">", ">", "^", "A", "A", "v", "A", "^", "<", "A", ">", "A", "A", "v", "A", "^", "A"}},
		{subSequence: []string{"<", "A", "A", "v", "<", "A", "A", ">", ">", "^", "A"}, move: []string{"v", "<", "<", "A", ">", ">", "^", "A", "A", "v", "<", "A", "<", "A", ">", ">", "^", "A", "A", "v", "A", "A", "^", "<", "A", ">", "A"}},
	}

	for _, tt := range tests {
		move := enterSubSequenceOnDPad(tt.subSequence)
		fmt.Println("Input:", tt.subSequence, "Result:", move)
		// fmt.Println("Expected:", tt.move)
		// fmt.Println("Actual:  ", move)
	
		if len(move) != len(tt.move) || !equalSlices(move, tt.move) {
			t.Errorf("enterSubSequenceOnDPad(%v) = %v; want %v", tt.subSequence, move, tt.move)
		}
	}
	
}

func TestEnterSequenceOnDPad(t *testing.T) {
	tests := []struct {
		sequence []string
		move     []string
	}{
		{sequence: []string{"^"}, move: []string{"v", "<", "<", "A", ">", ">", "^", "A"}},
		{sequence: []string{"<", "<", "^", "^", "A"}, move: []string{"v", "<", "A", "<", "A", "A", ">", ">", "^", "A", "A", "v", "A", "^", "<", "A", ">", "A", "A", "v", "A", "^", "A"}},
		{sequence: []string{"^", "^", "<", "<", "A"}, move: []string{"v", "<", "<", "A", ">", ">", "^", "A", "A", "v", "<", "A", "<", "A", ">", ">", "^", "A", "A", "v", "A", "A", "^", "<", "A", ">", "A"}},
	}

	for _, tt := range tests {

		commands := enterSequenceOnDPad(tt.sequence, 2)
		fmt.Println("Enter sequence", tt.sequence, "is", commands)

		if commands != len(tt.move) {
			t.Errorf("enterSequenceOnDPad(%v) = %v; want %v", tt.sequence, commands, tt.move)
		}
	}
}

func TestEnterCode(t *testing.T) {
	tests := []struct {
		sequence []string
		length   int
		cost     int
	}{
		{sequence: []string{"0", "2", "9", "A"}, length: 68, cost: 1972},
		{sequence: []string{"9", "8", "0", "A"}, length: 60, cost: 58800},
		{sequence: []string{"1", "7", "9", "A"}, length: 68, cost: 12172},
		{sequence: []string{"4", "5", "6", "A"}, length: 64, cost: 29184},
		{sequence: []string{"3", "7", "9", "A"}, length: 64, cost: 24256},
	}

	re := regexp.MustCompile(`\d+`)

	for _, tt := range tests {

		combinedStr := strings.Join(tt.sequence, "")
		digits := re.FindString(combinedStr)
		value, _ := strconv.Atoi(digits)

		sequenceLength := enterCodeOnNumberPad(tt.sequence, 2)

		fmt.Println("Enter sequence", tt.sequence, "length:", sequenceLength, "cost:", sequenceLength * value)

		if sequenceLength != tt.length {
			t.Errorf("enterCodeOnNumberPad(%v) = %v; want %v", tt.sequence, sequenceLength, tt.length)
		}
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
