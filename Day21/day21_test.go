package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestMoveOnNumPad_robot1(t *testing.T) {
	tests := []struct {
		start string
		end   string
		move  []string
	}{
		{start: "7", end: "9", move: []string{">", ">", "A"}},
		{start: "7", end: "5", move: []string{"v", ">", "A"}},
		{start: "7", end: "A", move: []string{"v", "v", "v", ">", ">", "A"}},
		{start: "A", end: "7", move: []string{"^", "^", "^", "<", "<", "A"}},
		{start: "A", end: "0", move: []string{"<", "A"}},
		{start: "3", end: "7", move: []string{"<", "<", "^", "^", "A"}},
		{start: "3", end: "7", move: []string{"<", "<", "^", "^", "A"}},
	}

	for _, tt := range tests {

		move := shortestRouteOnNumberPad(numPad[tt.start], numPad[tt.end])
		fmt.Println("Moving from", tt.start, "to", tt.end, "is", move)

		if len(move) != len(tt.move) {
			t.Errorf("moveOnNumberPad(%v, %v) = %v; want %v", tt.start, tt.end, move, tt.move)
		}
	}
}

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
	}

	for _, tt := range tests {

		move := enterSubSequenceOnDPad(tt.subSequence)
		fmt.Println("Enter sequence", tt.subSequence, "is", move)

		if len(move) != len(tt.move) {
			t.Errorf("enterSequenceOnDPad(%v) = %v; want %v", tt.subSequence, move, tt.move)
		}
	}
}

func TestEnterSequenceOnDPad(t *testing.T) {
	tests := []struct {
		sequence []string
		move     []string
	}{
		{sequence: []string{"v", "<", "<", "A", ">", ">", "^", "A"}, move: []string{"v", "<", "A", "<", "A", "A", "^", ">", ">", "A", "v", "A", "A", "^", "<", "A", ">", "A"}},
	}

	for _, tt := range tests {

		move := enterSequenceOnDPad(tt.sequence)
		fmt.Println("Enter sequence", tt.sequence, "is", move)

		if len(move) != len(tt.move) {
			t.Errorf("enterSequenceOnDPad(%v) = %v; want %v", tt.sequence, move, tt.move)
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

		move := enterCodeOnNumberPad(tt.sequence, 2)
		sequenceLength := len(move)

		fmt.Println("Enter sequence", tt.sequence, "length:", sequenceLength, "cost:", sequenceLength * value)

		if len(move) != tt.length {
			t.Errorf("enterCodeOnNumberPad(%v) = %v; want %v", tt.sequence, sequenceLength, tt.length)
		}
	}
}