package main

import (
	"testing"
)

func TestPositionAfterSeconds(t *testing.T) {

	xMax = 11
	yMax = 7
	xPos := 0
	yPos := 0

	//p=2,4 v=2,-3
	robot := robot{startingX: 2, startingY: 4, xDir: 2, yDir: -3}

	xPos, yPos = positionAfterSeconds(robot, 1)
	if xPos != 4 || yPos != 1 {
        t.Errorf("wrong position: x=%d, y=%d", xPos, yPos)
	}

	xPos, yPos = positionAfterSeconds(robot, 2)
	if xPos != 6 || yPos != 5 {
        t.Errorf("wrong position: x=%d, y=%d", xPos, yPos)
	}

	xPos, yPos = positionAfterSeconds(robot, 3)
	if xPos != 8 || yPos != 2 {
        t.Errorf("wrong position: x=%d, y=%d", xPos, yPos)
	}

	xPos, yPos = positionAfterSeconds(robot, 4)
	if xPos != 10 || yPos != 6 {
        t.Errorf("wrong position: x=%d, y=%d", xPos, yPos)
	}

	xPos, yPos = positionAfterSeconds(robot, 5)
	if xPos != 1 || yPos != 3 {
        t.Errorf("wrong position: x=%d, y=%d", xPos, yPos)
	}

}