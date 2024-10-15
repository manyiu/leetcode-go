package interleavingstring

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"a", "", "a", true},
		{"a", "", "b", false},
		{"", "a", "a", true},
		{"", "a", "b", false},
		{"a", "b", "a", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
	}
	for i, tt := range tests {
		t.Run(tt.s1+" "+tt.s2+" "+tt.s3, func(t *testing.T) {
			if got := isInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("Test %v, s1: %v, s2: %v, s3: %v, output = %v, want %v", i, tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}
