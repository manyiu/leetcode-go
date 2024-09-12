package shortestpathinbinarymatrix

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	grid := [][]int{
		{0, 1},
		{1, 0},
	}

	got := shortestPathBinaryMatrix(grid)
	want := 2

	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	grid = [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}

	got = shortestPathBinaryMatrix(grid)
	want = 4

	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	grid = [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}

	got = shortestPathBinaryMatrix(grid)
	want = -1

	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
