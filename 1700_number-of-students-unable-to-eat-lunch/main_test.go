package numberofstudentsunabletoeatlunch

import "testing"

func TestCountStudents(t *testing.T) {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	target := 0
	result := countStudents(students, sandwiches)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}

	students = []int{1, 1, 1, 0, 0, 1}
	sandwiches = []int{1, 0, 0, 0, 1, 1}
	target = 3
	result = countStudents(students, sandwiches)
	if target != result {
		t.Errorf("expect %d, got %d", target, result)
	}
}
