package containsduplicateii

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "example 2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "example 3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "example 4",
			nums: []int{99, 99},
			k:    2,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsNearbyDuplicate(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
