package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num      uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, test := range tests {
		if output := reverseBits(test.num); output != test.expected {
			t.Errorf("Expected reverseBits(%v) to be %v, but got %v", test.num, test.expected, output)
		}
	}
}
