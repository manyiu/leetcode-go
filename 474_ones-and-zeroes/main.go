package onesandzeroes

func countDigit(digits string) [2]int {
	zero := 0
	one := 0

	for _, digit := range digits {
		if digit == '0' {
			zero++
		} else if digit == '1' {
			one++
		}
	}

	return [2]int{zero, one}
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)

	for i := 0; i < len(dp); i++ {
		grid := make([][]int, m+1)

		for j := 0; j < len(grid); j++ {
			grid[j] = make([]int, n+1)
		}

		dp[i] = grid
	}

	for i := 1; i < len(dp); i++ {
		count := countDigit(strs[i-1])
		zeroCount := count[0]
		oneCount := count[1]

		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				if j-zeroCount >= 0 && k-oneCount >= 0 {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeroCount][k-oneCount]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}
