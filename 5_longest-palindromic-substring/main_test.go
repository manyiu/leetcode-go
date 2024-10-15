package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"bb", "bb"},
		{"ccc", "ccc"},
		{"aaaa", "aaaa"},
		{"abacdfgdcaba", "aba"},
		{"abcda", "a"},
	}

	for _, test := range tests {
		if got := longestPalindrome(test.s); got != test.expect {
			t.Errorf("longestPalindrome(%q) = %q; want %q", test.s, got, test.expect)
		}
	}
}
