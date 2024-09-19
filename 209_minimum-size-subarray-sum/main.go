package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1
	currSum := 0
	l, r := 0, 0

	for r < len(nums) {
		currSum += nums[r]

		for currSum >= target {
			if r-l+1 < min {
				min = r - l + 1
			}

			currSum -= nums[l]
			l++
		}

		r++
	}

	if min == len(nums)+1 {
		return 0
	}

	return min
}
