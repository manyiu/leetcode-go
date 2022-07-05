package maximumsubarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := 0

	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
