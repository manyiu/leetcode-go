package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	l, r := 0, 0

	for r < len(nums) {
		count := 1

		for r+1 < len(nums) && nums[r] == nums[r+1] {
			count++
			r++
		}

		for i := 0; i < 2 && i < count; i++ {
			nums[l] = nums[r]
			l++
		}

		r++
	}

	return l
}
