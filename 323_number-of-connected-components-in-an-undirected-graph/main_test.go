package numberofconnectedcomponentsinanundirectedgraph

import "testing"

func TestCountComponents(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	expected := 2
	result := countComponents(n, edges)
	if result != expected {
		t.Errorf("expected %d, but got %d", expected, result)
	}
}
