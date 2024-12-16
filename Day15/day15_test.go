package main

import (
	"testing"
)

func TestNextEmptyPosition(t *testing.T) {
	// Define the warehouse grid
	warehouse := [][]string{
		{"#", "#", "#", "#", "#"},
		{"#", ".", ".", "O", "#"},
		{"#", ".", "O", ".", "#"},
		{"#", "#", "#", "#", "#"},
	}

	// Define test cases
	tests := []struct {
		name          string
		start         position
		direction     position
		expected      position
	}{
		{
			name:      "Find next empty position to the right",
			start:     position{1, 1},
			direction: right,
			expected:  position{1, 2},
		},
		{
			name:      "Find next empty position skipping non-empty",
			start:     position{2, 1},
			direction: right,
			expected:  position{2, 3},
		},
		{
			name:      "Find no space when moving right",
			start:     position{1, 2},
			direction: right,
			expected:  position{-1, -1},
		},
		{
			name:      "Find next empty position moving down",
			start:     position{1, 1},
			direction: down,
			expected:  position{2, 1},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nextEmptyPosition(tt.start, tt.direction, warehouse)
			if result != tt.expected {
				t.Errorf("nextEmptyPosition(%v, %v) = %v; want %v", tt.start, tt.direction, result, tt.expected)
			}
		})
	}
}


func TestMoveBigBox(t *testing.T) {

	// Define test cases
	tests := []struct {
		name          		string
		start         		position
		startingWarehouse 	[][]string
		direction     		position
		expected      		[][]string
	}{
		// {
		// 	name:      "Move big box right",
		// 	start:     position{0, 1},
		// 	startingWarehouse: [][]string{
		// 		{"#", "@", "[", "]", ".", "#"},
		// 	},
		// 	direction: right,
		// 	expected:  [][]string{
		// 		{"#", ".", "@", "[", "]", "#"},
		// 	},
		// },
		// {
		// 	name:      "Try move big box right and hits the wall",
		// 	start:     position{0, 2},
		// 	startingWarehouse: [][]string{
		// 		{"#", ".", "@", "[", "]", "#"},
		// 	},
		// 	direction: right,
		// 	expected:  [][]string{
		// 		{"#", ".", "@", "[", "]", "#"},
		// 	},
		// },
		// {
		// 	name:      "Move big box left",
		// 	start:     position{0, 5},
		// 	startingWarehouse: [][]string{
		// 		{"#", ".", ".", "[", "]", "@"},
		// 	},
		// 	direction: left,
		// 	expected:  [][]string{
		// 		{"#", ".", "[", "]", "@", "."},
		// 	},
		// },
		// {
		// 	name:      "Move 2 big boxes right",
		// 	start:     position{0, 1},
		// 	startingWarehouse: [][]string{
		// 		{"#", "@", "[", "]", "[", "]", ".", "."},
		// 	},
		// 	direction: right,
		// 	expected:  [][]string{
		// 		{"#", ".", "@", "[", "]", "[", "]", "."},
		// 	},
		// },
		// {
		// 	name:      "Move 1 big boxes up",
		// 	start:     position{3, 3},
		// 	startingWarehouse: [][]string{
		// 		{"#", "#", "#", "#", "#", "#"},
		// 		{"#", ".", ".", ".", ".", "#"},
		// 		{"#", ".", "[", "]", ".", "#"},
		// 		{"#", ".", ".", "@", ".", "#"},
		// 		{"#", "#", "#", "#", "#", "#"},
		// 	},
		// 	direction: up,
		// 	expected:  [][]string{
		// 		{"#", "#", "#", "#", "#", "#"},
		// 		{"#", ".", "[", "]", ".", "#"},
		// 		{"#", ".", ".", "@", ".", "#"},
		// 		{"#", ".", ".", ".", ".", "#"},
		// 		{"#", "#", "#", "#", "#", "#"},
		// 	},
		// },
		{
			name:      "Move 2 big boxes up",
			start:     position{4, 3},
			startingWarehouse: [][]string{
				{"#", "#", "#", "#", "#", "#"},
				{"#", ".", ".", ".", ".", "#"},
				{"#", "[", "]", "[", "]", "#"},
				{"#", ".", "[", "]", ".", "#"},
				{"#", ".", ".", "@", ".", "#"},
				{"#", "#", "#", "#", "#", "#"},
			},
			direction: up,
			expected:  [][]string{
				{"#", "#", "#", "#", "#", "#"},
				{"#", "[", "]", "[", "]", "#"},
				{"#", ".", "[", "]", ".", "#"},
				{"#", ".", ".", "@", ".", "#"},
				{"#", ".", ".", ".", ".", "#"},
				{"#", "#", "#", "#", "#", "#"},
			},
		},
		{
			name:      "Move 2 diagonal big boxes up",
			start:     position{4, 3},
			startingWarehouse: [][]string{
				{"#", "#", "#", "#", "#", "#"},
				{"#", ".", ".", ".", ".", "#"},
				{"#", "[", "]", ".", ".", "#"},
				{"#", ".", "[", "]", ".", "#"},
				{"#", ".", ".", "@", ".", "#"},
				{"#", "#", "#", "#", "#", "#"},
			},
			direction: up,
			expected:  [][]string{
				{"#", "#", "#", "#", "#", "#"},
				{"#", "[", "]", ".", ".", "#"},
				{"#", ".", "[", "]", ".", "#"},
				{"#", ".", ".", "@", ".", "#"},
				{"#", ".", ".", ".", ".", "#"},
				{"#", "#", "#", "#", "#", "#"},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			warehouse := tt.startingWarehouse
			moveBigBox(&tt.start, tt.direction, warehouse)
			printWarehouse(warehouse)

			for r := range warehouse {
				for c := range warehouse[0] {
					if warehouse[r][c] != tt.expected[r][c] {
						t.Errorf("warehouse[%v, %v]:%v != expected[%v, %v]:%v", r, c, warehouse[r][c], r, c, tt.expected[r][c])
					}
				}
			}
		})
	}
}
