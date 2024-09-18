package maximumsumcircularsubarray

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)

	for n == 0 {
		return 0
	}

	min, max := nums[0], nums[0]
	currMin, currMax := 0, 0
	total := 0

	for i := 0; i < len(nums); i++ {
		if currMin > 0 {
			currMin = nums[i]
		} else {
			currMin += nums[i]
		}

		if currMax < 0 {
			currMax = nums[i]
		} else {
			currMax += nums[i]
		}

		total += nums[i]

		if currMin < min {
			min = currMin
		}

		if currMax > max {
			max = currMax
		}
	}

	if max < 0 {
		return max
	}

	if total-min > max {
		return total - min
	}

	return max
}
