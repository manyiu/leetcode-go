package productofarrayexceptself

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	actual := productExceptSelf(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}
