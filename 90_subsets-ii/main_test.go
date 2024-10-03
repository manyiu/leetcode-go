package subsetsii

import "testing"

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 2},
			want: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
	}

	for _, tt := range tests {
		got := subsetsWithDup(tt.nums)

		if len(got) != len(tt.want) {
			t.Fatalf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.want)
		}

		for i := 0; i < len(got); i++ {
			if len(got[i]) != len(tt.want[i]) {
				t.Fatalf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.want)
			}

			for j := 0; j < len(got[i]); j++ {
				if got[i][j] != tt.want[i][j] {
					t.Fatalf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.want)
				}
			}
		}
	}
}
