package main

import (
	"fmt"
	"testing"
)

func TestMix(t *testing.T) {
	tests := []struct {
		input    	int
		secret		int
		expected   	int
	}{
		{input: 15, secret: 42, expected: 37},
	}

	for _, tt := range tests {

		result := mix(tt.input, tt.secret)

		if result != tt.expected {
			t.Errorf("mix(%v, %v) = %v; want %v", tt.input, tt.secret, result, tt.expected)
		}
	}
}

func TestPrine(t *testing.T) {
	tests := []struct {
		secret		int
		expected   	int
	}{
		{secret: 100000000, expected: 16113920},
	}

	for _, tt := range tests {

		result := prune(tt.secret)

		if result != tt.expected {
			t.Errorf("prune(%v) = %v; want %v", tt.secret, result, tt.expected)
		}
	}
}

func TestProcessNextSecret(t *testing.T) {
	tests := []struct {
		secret		int
		expected   	int
	}{
		{secret: 123, expected: 15887950},
		{secret: 15887950, expected: 16495136},
		{secret: 16495136, expected: 527345},
		{secret: 527345, expected: 704524},
		{secret: 704524, expected: 1553684},
		{secret: 1553684, expected: 12683156},
		{secret: 12683156, expected: 11100544},
		{secret: 11100544, expected: 12249484},
		{secret: 12249484, expected: 7753432},
		{secret: 7753432, expected: 5908254},
	}

	for _, tt := range tests {

		result := nextSecretNumber(tt.secret)

		if result != tt.expected {
			t.Errorf("nextSecretNumber(%v) = %v; want %v", tt.secret, result, tt.expected)
		}
	}
}

func TestNthSecret(t *testing.T) {
	tests := []struct {
		secret		int
		expected   	int
	}{
		{secret: 1, expected: 8685429},
		{secret: 10, expected: 4700978},
		{secret: 100, expected: 15273692},
		{secret: 2024, expected: 8667524},
	}

	for _, tt := range tests {

		result := nthSecretNumber(tt.secret, 2000)

		if result != tt.expected {
			t.Errorf("nextSecretNumber(%v) = %v; want %v", tt.secret, result, tt.expected)
		}
	}
}

func TestFindSequenceOfChanges(t *testing.T) {
    tests := []struct {
        secret   int
		sequence []int
        expected int
    }{
        {
            secret: 123,
			sequence: []int{-1,-1,0,2},
            expected: 6,
        },
        {
            secret: 1,
			sequence: []int{-2,1,-1,3},
            expected: 7,
        },
    }

    for _, test := range tests {
        result := findSequenceOfChanges(test.secret)

        if result[fmt.Sprintf("%v", test.sequence)] != test.expected {
            t.Errorf("findSequenceOfChanges(%d) was %v; want %v", test.secret, result[fmt.Sprintf("%v", test.sequence)], test.expected)
        }
    }
}