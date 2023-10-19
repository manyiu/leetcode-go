package rottingoranges

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	expected := 4

	actual := orangesRotting(grid)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
