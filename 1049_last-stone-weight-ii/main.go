package laststoneweightii

func lastStoneWeightII(stones []int) int {
	sum := 0

	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2

	dp := make([][]int, len(stones)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j-stones[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j-stones[i-1]]+stones[i-1], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return sum - (2 * dp[len(stones)][target])
}
