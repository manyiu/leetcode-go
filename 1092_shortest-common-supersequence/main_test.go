package shortestcommonsupersequence

import "testing"

func TestShortestCommonSupersequence(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"abac", "cab", "cabac"},
		{"aaaaaaaa", "aaaaaaaa", "aaaaaaaa"},
		{"abc", "def", "abcdef"},
		{"geek", "eke", "geeke"},
	}
	for i, tt := range tests {
		t.Run(tt.str1+" "+tt.str2, func(t *testing.T) {
			if got := shortestCommonSupersequence(tt.str1, tt.str2); got != tt.want {
				t.Errorf("Test %v, str1: %v, str2: %v, output = %v, want %v", i, tt.str1, tt.str2, got, tt.want)
			}
		})
	}
}
