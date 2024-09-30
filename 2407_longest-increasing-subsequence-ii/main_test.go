package longestincreasingsubsequenceii

import "testing"

func TestLengthOfLIS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"testcase1", []int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3, 5},
		{"testcase2", []int{7, 4, 5, 1, 8, 12, 4, 7}, 5, 4},
		{"testcase3", []int{1, 5}, 1, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := lengthOfLIS(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
