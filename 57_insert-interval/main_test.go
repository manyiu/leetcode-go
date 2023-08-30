package insertinterval

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals [][]int
		new       []int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			new:       []int{2, 5},
			want:      [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			new:       []int{4, 8},
			want:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals: [][]int{},
			new:       []int{5, 7},
			want:      [][]int{{5, 7}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{2, 3},
			want:      [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{2, 7},
			want:      [][]int{{1, 7}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{6, 8},
			want:      [][]int{{1, 5}, {6, 8}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{0, 3},
			want:      [][]int{{0, 5}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{0, 0},
			want:      [][]int{{0, 0}, {1, 5}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{0, 6},
			want:      [][]int{{0, 6}},
		},
		{
			intervals: [][]int{{1, 5}},
			new:       []int{0, 8},
			want:      [][]int{{0, 8}},
		},
	}

	for _, tt := range tests {
		got := insert(tt.intervals, tt.new)
		if len(got) != len(tt.want) {
			t.Errorf("insert(%v, %v) got %v, want %v", tt.intervals, tt.new, got, tt.want)
		}
		for i := 0; i < len(got); i++ {
			if len(got[i]) != len(tt.want[i]) {
				t.Errorf("insert(%v, %v) got %v, want %v", tt.intervals, tt.new, got, tt.want)
			}
			for j := 0; j < len(got[i]); j++ {
				if got[i][j] != tt.want[i][j] {
					t.Errorf("insert(%v, %v) got %v, want %v", tt.intervals, tt.new, got, tt.want)
				}
			}
		}
	}
}
