package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	prefixSumCount := map[int]int{
		0: 1,
	}
	prefixSum := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]

		if _, ok := prefixSumCount[prefixSum-k]; ok {
			count += prefixSumCount[prefixSum-k]
		}

		prefixSumCount[prefixSum]++
	}

	return count
}
