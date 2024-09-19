package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	l, r := 0, 1

	for r < len(nums) {
		if nums[l] == nums[r] {
			if r-l >= 2 {
				nums = append(nums[:r], nums[r+1:]...)
			} else {
				r++
			}
		} else {
			l = r
			r++
		}
	}

	return len(nums)
}
