package courseschedule

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	input := [][]int{
		{1, 0},
		{0, 1},
	}
	expected := false
	actual := canFinish(2, input)

	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
