package coursescheduleii

import "testing"

func isSameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			2,
			[][]int{{1, 0}},
			[]int{0, 1},
		},
		{
			4,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[]int{0, 1, 2, 3},
		},
		{
			2,
			[][]int{{1, 0}, {0, 1}},
			[]int{},
		},
	}
	for _, test := range tests {
		if got := findOrder(test.numCourses, test.prerequisites); !isSameSlice(got, test.want) {
			t.Errorf("findOrder(%v, %v) = %v, want %v", test.numCourses, test.prerequisites, got, test.want)
		}
	}
}
