package lettercombinationsofaphonenumber

import "testing"

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		expect []string
	}{
		{
			name:   "test1",
			digits: "23",
			expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "test2",
			digits: "",
			expect: []string{},
		},
		{
			name:   "test3",
			digits: "2",
			expect: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		if got := letterCombinations(test.digits); !equal(got, test.expect) {
			t.Errorf("letterCombinations(%s) = %v; want %v", test.digits, got, test.expect)
		}
	}
}
