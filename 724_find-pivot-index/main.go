package findpivotindex

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	for j := 0; j < len(nums); j++ {
		if j == 0 {
			if nums[len(nums)-1] == nums[j] {
				return j
			}
		} else {
			if nums[len(nums)-1]-nums[j-1] == nums[j] {
				return j
			}
		}
	}

	return -1
}
