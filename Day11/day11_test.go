package main

import (
	"testing"
)

func TestNumberOfDigits(t *testing.T) {

	if numberOfDigits(12) != 2 {
		t.Errorf("wrong number of digits")
	}
	if numberOfDigits(123456) != 6 {
		t.Errorf("wrong number of digits")
	}
	if numberOfDigits(1) != 1 {
		t.Errorf("wrong number of digits")
	}

}

func TestBlinkAtStone(t *testing.T) {

	stones := blinkAtStone(0)
	if len(stones) != 1 || stones[0] != 1 {
		t.Errorf("issue with blink at 0")
	}

	stones = blinkAtStone(1234) 
	if len(stones) != 2 || stones[0] != 12 || stones[1] != 34{
		t.Errorf("issue with blink at 1234")
	}

	stones = blinkAtStone(1) 
	if stones[0] != 2024 {
		t.Errorf("issue with blink at 1")
	}

}