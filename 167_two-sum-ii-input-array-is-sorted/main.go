package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for numbers[l]+numbers[r] != target {
		s := numbers[l] + numbers[r]

		if s > target {
			r--
		} else {
			l++
		}
	}

	return []int{l + 1, r + 1}
}
