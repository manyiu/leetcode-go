package longestrepeatingcharacterreplacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AABABBA", 0, 2},
		{"AABABBA", 3, 7},
		{"AABABBA", 4, 7},
		{"AABABBA", 5, 7},
		{"AABABBA", 6, 7},
		{"AABABBA", 7, 7},
		{"AABABBA", 8, 7},
		{"AABABBA", 9, 7},
		{"AABABBA", 10, 7},
		{"AABABBA", 11, 7},
		{"AABABBA", 12, 7},
		{"AABABBA", 13, 7},
		{"AABABBA", 14, 7},
		{"AABABBA", 15, 7},
		{"AABABBA", 16, 7},
		{"AABABBA", 17, 7},
		{"AABABBA", 18, 7},
		{"AABABBA", 19, 7},
		{"AABABBA", 20, 7},
		{"AABABBA", 21, 7},
		{"AABABBA", 22, 7},
		{"AABABBA", 23, 7},
		{"AABABBA", 24, 7},
		{"AABABBA", 25, 7},
		{"AABABBA", 26, 7},
		{"AABABBA", 27, 7},
		{"AABABBA", 28, 7},
		{"AABABBA", 29, 7},
		{"AABABBA", 30, 7},
		{"AABABBA", 31, 7},
		{"AABABBA", 32, 7},
		{"AABABBA", 33, 7},
		{"ABAA", 0, 2},
	}

	for _, test := range tests {
		if got := characterReplacement(test.s, test.k); got != test.output {
			t.Errorf("characterReplacement(%v, %v) = %v; want %v\n", test.s, test.k, got, test.output)
		}
	}
}
