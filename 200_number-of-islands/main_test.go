package numberofislands

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	input := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}

	expected := 1
	actual := numIslands(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
