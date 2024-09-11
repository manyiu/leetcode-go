package maxareaofisland

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	expected := 6
	if res := maxAreaOfIsland(grid); res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}

	grid = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	expected = 0
	if res := maxAreaOfIsland(grid); res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
