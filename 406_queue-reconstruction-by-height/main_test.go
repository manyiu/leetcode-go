package queuereconstructionbyheight

import "testing"

func TestReconstructQueue(t *testing.T) {
	people := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	expected := [][]int{
		{5, 0},
		{7, 0},
		{5, 2},
		{6, 1},
		{4, 4},
		{7, 1},
	}
	result := reconstructQueue(people)
	if len(result) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
	for i := range result {
		if result[i][0] != expected[i][0] || result[i][1] != expected[i][1] {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}
