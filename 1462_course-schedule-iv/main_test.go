package coursescheduleiv

import "testing"

func isSameSlice(a, b []bool) bool {
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

func TestCheckIfPrerequisite(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		want          []bool
	}{
		{
			2,
			[][]int{{1, 0}},
			[][]int{{0, 1}, {1, 0}},
			[]bool{false, true},
		},
		{
			2,
			[][]int{},
			[][]int{{1, 0}, {0, 1}},
			[]bool{false, false},
		},
		{
			3,
			[][]int{{1, 2}, {1, 0}, {2, 0}},
			[][]int{{1, 0}, {1, 2}},
			[]bool{true, true},
		},
	}
	for i, test := range tests {
		if got := checkIfPrerequisite(test.numCourses, test.prerequisites, test.queries); !isSameSlice(got, test.want) {
			t.Errorf("Test %v: checkIfPrerequisite(%v, %v, %v) = %v, want %v", i, test.numCourses, test.prerequisites, test.queries, got, test.want)
		}
	}
}
