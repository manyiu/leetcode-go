package partitionequalsubsetsum

func canPartition(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	dp := make([][]bool, len(nums)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
	}

	dp[0][0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	result := dp[len(nums)][target]

	return result
}
