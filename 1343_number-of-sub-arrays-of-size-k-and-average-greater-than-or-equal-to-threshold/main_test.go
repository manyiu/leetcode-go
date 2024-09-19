package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	tests := []struct {
		arr       []int
		k         int
		threshold int
		want      int
	}{
		{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4, 3},
		{[]int{1, 1, 1, 1, 1}, 1, 0, 5},
		{[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5, 6},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 7, 7, 1},
		{[]int{4, 4, 4, 4}, 4, 1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := numOfSubarrays(tt.arr, tt.k, tt.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
