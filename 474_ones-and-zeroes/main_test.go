package onesandzeroes

import "testing"

func TestFindMaxForm(t *testing.T) {
	tests := []struct {
		strs []string
		m    int
		n    int
		want int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findMaxForm(tt.strs, tt.m, tt.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
