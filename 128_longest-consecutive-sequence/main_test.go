package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"two elements", []int{1, 2}, 2},
		{"three elements", []int{100, 4, 200}, 1},
		{"four elements", []int{100, 4, 200, 1}, 1},
		{"five elements", []int{100, 4, 200, 1, 3}, 2},
		{"six elements", []int{100, 4, 200, 1, 3, 2}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestConsecutive(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums: %v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
