package distinctsubsequences

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
	}
	for _, tt := range tests {
		t.Run(tt.s+" "+tt.t, func(t *testing.T) {
			if got := numDistinct(tt.s, tt.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
