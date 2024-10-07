package swiminrisingwater

import "testing"

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			name: "test 1",
			grid: [][]int{
				{0, 2},
				{1, 3},
			},
			expect: 3,
		},
		{
			name: "test 2",
			grid: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expect: 8,
		},
		{
			name: "test 3",
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			expect: 16,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := swimInWater(test.grid)
			if got != test.expect {
				t.Errorf("expect %d, got %d", test.expect, got)
			}
		})
	}
}
