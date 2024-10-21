package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	st := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect float64
	}{
		{"test 1", []int{1, 3}, []int{2}, 2.0},
		{"test 2", []int{1, 2}, []int{3, 4}, 2.5},
		{"test 3", []int{}, []int{1}, 1.0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findMedianSortedArrays(tt.nums1, tt.nums2)
			if out != tt.expect {
				t.Fatalf("with input nums1:%v and nums2:%v wanted %.2f but got %.2f", tt.nums1, tt.nums2, tt.expect, out)
			}
			t.Log("pass")
		})
	}
}
