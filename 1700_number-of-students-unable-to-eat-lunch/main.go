package numberofstudentsunabletoeatlunch

func countStudents(students []int, sandwiches []int) int {
	studentCount := map[int]int{
		0: 0,
		1: 0,
	}

	for _, student := range students {
		studentCount[student]++
	}

	for _, sandwich := range sandwiches {
		if studentCount[sandwich] == 0 {
			break
		}

		studentCount[sandwich]--
	}

	return studentCount[0] + studentCount[1]
}
