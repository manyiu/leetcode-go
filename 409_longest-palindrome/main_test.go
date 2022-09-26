package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s string
		l int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
	}
	for _, test := range tests {
		if l := longestPalindrome(test.s); l != test.l {
			t.Errorf("longestPalindrome(%s) = %d; want %d", test.s, l, test.l)
		}
	}
}
