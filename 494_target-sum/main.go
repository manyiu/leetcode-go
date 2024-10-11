package targetsum

func findTargetSumWays(nums []int, target int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if (sum+target)%2 != 0 || (sum+target) < 0 {
		return 0
	}

	sumTarget := (sum + target) / 2

	dp := make([][]int, len(nums)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sumTarget+1)
	}

	dp[0][0] = 1

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][sumTarget]
}
