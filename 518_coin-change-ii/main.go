package coinchangeii

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := 1; j < len(dp); j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}

	return dp[amount]
}
