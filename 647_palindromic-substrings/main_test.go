package palindromicsubstrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"abc", 3},
		{"aaa", 6},
		{"abba", 6},
		{"a", 1},
		{"", 0},
		{"fdsklf", 6},
	}
	for _, test := range tests {
		if got := countSubstrings(test.in); got != test.want {
			t.Errorf("countSubstrings(%q) = %v; want %v\n", test.in, got, test.want)
		}
	}
}
