package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	target := 1

	if output := lastStoneWeight(stones); output != target {
		t.Errorf("Expected %d but received %d", target, output)
	}

	stones = []int{2, 2}
	target = 0

	if output := lastStoneWeight(stones); output != target {
		t.Errorf("Expected %d but received %d", target, output)
	}

	stones = []int{1, 3}
	target = 2
	if output := lastStoneWeight(stones); output != target {
		t.Errorf("Expected %d but received %d", target, output)
	}
}
