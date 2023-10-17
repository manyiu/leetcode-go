package coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	input := []int{1, 2, 5}
	expected := 3
	actual := coinChange(input, 11)

	if actual != expected {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
