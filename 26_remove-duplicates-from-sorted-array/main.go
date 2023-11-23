package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i, j int

	for j = 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
