package main

import "testing"

func TestMove(t *testing.T) {

	sp := Position{row: 0, column: 0}

	up := Position{row: -1, column: 0}
	down := Position{row: 1, column: 0}
	left := Position{row: 0, column: -1}
	right := Position{row: 0, column: 1}

	if move(sp, Up) != up {
		t.Errorf("Up failed")
	}
	if move(sp, Down) != down {
		t.Errorf("Up failed")
	}
	if move(sp, Left) != left {
		t.Errorf("Up failed")
	}
	if move(sp, Right) != right {
		t.Errorf("Up failed")
	}

}