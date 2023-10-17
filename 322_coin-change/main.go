package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for j := 1; j < len(dp); j++ {
		for _, coin := range coins {
			if j-coin >= 0 {
				cur := dp[j]
				can := dp[j-coin] + 1

				if can < cur {
					dp[j] = can
				}
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
