package ransomnote

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote, magazine string
		want                 bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, tt := range tests {
		t.Run(tt.ransomNote, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
