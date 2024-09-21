package subarraysumequalsk

import "testing"

func TestSubarraySum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		// {"empty slice", []int{}, 1, 0},
		// {"single element", []int{1}, 1, 1},
		// {"two elements", []int{1, 2}, 3, 1},
		// {"testcase1", []int{1, 1, 1}, 2, 2},
		// {"testcase2", []int{1, 2, 3}, 3, 2},
		// {"testcase3", []int{1, 2, 3, 4}, 3, 2},
		{"testcase4", []int{1, 2, 3, 4}, 5, 2},
		// {"testcase5", []int{1, 2, 3, 4}, 6, 1},
		// {"testcase6", []int{1, 2, 3, 4}, 7, 1},
		// {"testcase7", []int{1, 2, 3, 4}, 8, 1},
		// {"testcase8", []int{1, 2, 3, 4}, 9, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := subarraySum(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("%v with input nums: %v and k: %d wanted %d but got %d", tt.name, tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
