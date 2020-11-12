package core

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	got := NewGrid(3, 3)
	if (got.Rows != 3 || got.Columns != 3) {
		t.Errorf("Expected (3,3) dimensions, but got (%d, %d)", got.Rows, got.Columns)
	}
}

func TestEachRow(t *testing.T) {
	grid := NewGrid(3, 3)
	count := 0
	cb := func(row []*Cell, _ int) {
		count += 1
	}
	grid.EachRow(cb)
	if (count != 3) {
		t.Errorf("Expected cb to have been called 3 times, got %v", count)
	}
}

func TestLinks(t *testing.T) {
	sut := NewGrid(2, 2)
	grid := sut.grid

	topLeft := grid[0][0]
	if (topLeft.West != nil || topLeft.North != nil) {
		t.Errorf("Expected top-left corner to have no neighbours to W,N")
	}
	if (topLeft.East == nil || topLeft.South == nil) {
		t.Errorf("Expected top-left corner to have neighbours to S,E")
	}

	topRight := grid[0][1]
	if (topRight.East != nil || topRight.North != nil) {
		t.Errorf("Expected top-right corner to have no neighbours to E,N")
	}
	if (topRight.West == nil || topRight.South == nil) {
		t.Errorf("Expected top-right corner to have neighbours to S,W")
	}

	bottomLeft := grid[1][0]
	if (bottomLeft.West != nil || bottomLeft.South != nil) {
		t.Errorf("Expected bottom-left corner to have no neighbours to W,S")
	}
	if (bottomLeft.East == nil || bottomLeft.North == nil) {
		t.Errorf("Expected bottom-left corner to have neighbours to S,W")
	}

	bottomRight := grid[1][1]
	if (bottomRight.East != nil || bottomRight.South != nil) {
		t.Errorf("Expected bottom-right corner to have no neighbours to E,S")
	}
}
