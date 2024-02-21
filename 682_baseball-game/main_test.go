package baseballgame

import "testing"

func TestCalPoints(t *testing.T) {
	operations := []string{"5", "2", "C", "D", "+"}
	expected := 30
	actual := calPoints(operations)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	operations = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	expected = 27
	actual = calPoints(operations)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}

	operations = []string{"1", "C"}
	expected = 0
	actual = calPoints(operations)
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
