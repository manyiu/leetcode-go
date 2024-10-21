package repeatedstringmatch

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	st := []struct {
		name string
		a, b string
		exp  int
	}{
		{"testcase1", "abcd", "cdabcdab", 3},
		{"testcase2", "a", "aa", 2},
		{"testcase3", "a", "a", 1},
		{"testcase4", "abc", "wxyz", -1},
		{"testcase5", "abc", "cabcabca", 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := repeatedStringMatch(tt.a, tt.b)
			if out != tt.exp {
				t.Fatalf("with input a:%s and b:%s wanted %d but got %d", tt.a, tt.b, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
