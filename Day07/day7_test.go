package main

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		// Test Case 1: Single element
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		// Test Case 2: Two elements
		{
			name:     "Two elements",
			input:    []int{1, 2},
			expected: []int{3, 2, 12}, // Possible results: 1+2, 1*2, concatenation
		},
		// Test Case 3: Three elements
		{
			name:     "Three elements",
			input:    []int{1, 2, 3},
			expected: []int{6, 9, 33, 5, 6, 23, 15, 36, 123}, // Simplified result from recursion
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combine(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("combine(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	test := []int{1,2,3,4,5}
	if !contains(test, 1) {
		t.Fatalf("Contains test failed")
	}
}

