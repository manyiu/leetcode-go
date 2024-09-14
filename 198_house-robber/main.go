package houserobber

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	nums[2] = nums[0] + nums[2]

	for i := 3; i < len(nums); i++ {
		nums[i] = nums[i] + max(nums[i-2], nums[i-3])
	}

	return max(nums[len(nums)-1], nums[len(nums)-2])
}
