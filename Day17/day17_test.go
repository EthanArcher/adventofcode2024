package main

import (
	"testing"
)

func TestAdv(t *testing.T) {
	tests := []struct {
		initialA int
		operand  int
		expected int
	}{
		{initialA: 8, operand: 1, expected: 4},
		{initialA: 8, operand: 2, expected: 2},
		{initialA: 8, operand: 3, expected: 1},
		{initialA: 8, operand: 4, expected: 0},
		{initialA: 15, operand: 1, expected: 7},
		{initialA: 15, operand: 2, expected: 3},
		{initialA: 15, operand: 3, expected: 1},
		{initialA: 15, operand: 4, expected: 0},
	}

	for _, tt := range tests {
		A = tt.initialA
		adv(tt.operand)
		if A != tt.expected {
			t.Errorf("adv(%d) = %d; want %d", tt.operand, A, tt.expected)
		}
	}
}
func TestBxl(t *testing.T) {
	tests := []struct {
		initialB int
		operand  int
		expected int
	}{
		{initialB: 0, operand: 1, expected: 1},
		{initialB: 1, operand: 1, expected: 0},
		{initialB: 2, operand: 3, expected: 1},
		{initialB: 4, operand: 4, expected: 0},
		{initialB: 5, operand: 2, expected: 7},
	}

	for _, tt := range tests {
		B = tt.initialB
		bxl(tt.operand)
		if B != tt.expected {
			t.Errorf("bxl(%d) = %d; want %d", tt.operand, B, tt.expected)
		}
	}
}

func TestBst(t *testing.T) {
	tests := []struct {
		operand  int
		expected int
	}{
		{operand: 1, expected: 1},
		{operand: 0, expected: 0},
		{operand: 4, expected: 5}, // A
		{operand: 5, expected: 0}, // B
		{operand: 6, expected: 1}, // C
	}

	
	for _, tt := range tests {
		A = 5
		B = 8
		C = 9
		bst(tt.operand)
		if B != tt.expected {
			t.Errorf("bst(%d) = %d; want %d", tt.operand, B, tt.expected)
		}
	}
}

func TestJnz(t *testing.T) {
	tests := []struct {
		initialA int
		p		 int
		operand  int
		expected int
	}{
		{initialA: 0, p:0, operand: 1, expected: 2},
		{initialA: 0, p:4, operand: 5, expected: 6},
		{initialA: 1, p:4, operand: 1, expected: 1},
		{initialA: 8, p:6, operand: 5, expected: 5},
	}

	for _, tt := range tests {
		A = tt.initialA
		pointerPosition = tt.p
		jnz(tt.operand)

		if pointerPosition != tt.expected {
			t.Errorf("jnz(%d) = %d; want %d", tt.operand, pointerPosition, tt.expected)
		}
	}
}

func TestBxc(t *testing.T) {
	tests := []struct {
		initialB int
		initialC int
		expected int
	}{
		{initialB: 0, initialC: 1, expected: 1},
		{initialB: 1, initialC: 1, expected: 0},
		{initialB: 2, initialC: 3, expected: 1},
		{initialB: 4, initialC: 4, expected: 0},
		{initialB: 5, initialC: 2, expected: 7},
	}

	for _, tt := range tests {
		B = tt.initialB
		C = tt.initialC
		bxc(0)
		if B != tt.expected {
			t.Errorf("bxc() = %d; want %d", B, tt.expected)
		}
	}
}

func TestOut(t *testing.T) {
	tests := []struct {
		operand  int
		expected string
	}{
		{operand: 1, expected: "1,"},
		{operand: 0, expected: "0,"},
		{operand: 4, expected: "5,"}, // A
		{operand: 5, expected: "0,"}, // B
		{operand: 6, expected: "1,"}, // C
	}

	for _, tt := range tests {
		output.Reset()
		A = 5
		B = 8
		C = 9
		out(tt.operand)
		if output.String() != tt.expected {
			t.Errorf("out(%d) = %v; want %v", tt.operand, output.String(), tt.expected)
		}
	}
}

func TestRunProgram(t *testing.T) {
	tests := []struct {
		initialA int
		initialB int
		initialC int
		program  []int
		expectedA int
		expectedB int
		expectedC int
		expectedOutput string
	}{
		{initialA: 0, initialB: 0, initialC: 9, program: []int{2,6}, expectedA: 0, expectedB: 1, expectedC: 9, expectedOutput: ""},
		{initialA: 10, initialB: 0, initialC: 0, program: []int{5,0,5,1,5,4}, expectedA: 10, expectedB: 0, expectedC: 0, expectedOutput: "0,1,2,"},
		{initialA: 2024, initialB: 0, initialC: 0, program: []int{0,1,5,4,3,0}, expectedA: 0, expectedB: 0, expectedC: 0, expectedOutput: "4,2,5,6,7,7,7,7,3,1,0,"},
		{initialA: 0, initialB: 29, initialC: 0, program: []int{1,7}, expectedA: 0, expectedB: 26, expectedC: 0, expectedOutput: ""},
		{initialA: 0, initialB: 2024, initialC: 43690, program: []int{4,0}, expectedA: 0, expectedB: 44354, expectedC: 43690, expectedOutput: ""},
	}

	for _, tt := range tests {
		A = tt.initialA
		B = tt.initialB
		C = tt.initialC
		output.Reset()
		pointerPosition = 0
		runProgram(tt.program)
		if A != tt.expectedA {
			t.Errorf("runProgram(%v) = %d; want %d", tt.program, A, tt.expectedA)
		}
		if B != tt.expectedB {
			t.Errorf("runProgram(%v) = %d; want %d", tt.program, B, tt.expectedB)
		}
		if C != tt.expectedC {
			t.Errorf("runProgram(%v) = %d; want %d", tt.program, C, tt.expectedC)
		}
		if output.String() != tt.expectedOutput {
			t.Errorf("runProgram(%v) = %v; want %v", tt.program, output.String(), tt.expectedOutput)
		}
	}
}
