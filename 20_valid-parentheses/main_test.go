package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
	}
	for _, test := range tests {
		if got := isValid(test.s); got != test.want {
			t.Errorf("isValid(%q) = %v", test.s, got)
		}
	}
}
