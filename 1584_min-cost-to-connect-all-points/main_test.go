package mincosttoconnectallpoints

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		expect int
	}{
		{
			name: "test 1",
			points: [][]int{
				{0, 0},
				{2, 2},
				{3, 10},
				{5, 2},
				{7, 0},
			},
			expect: 20,
		},
		{
			name: "test 2",
			points: [][]int{
				{3, 12},
				{-2, 5},
				{-4, 1},
			},
			expect: 18,
		},
		{
			name: "test 3",
			points: [][]int{
				{0, 0},
				{1, 1},
				{1, 0},
				{-1, 1},
			},
			expect: 4,
		},
		{
			name: "test 4",
			points: [][]int{
				{-1000000, -1000000},
				{1000000, 1000000},
			},
			expect: 4000000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minCostConnectPoints(test.points)
			if got != test.expect {
				t.Errorf("expect %d, got %d", test.expect, got)
			}
		})
	}
}
