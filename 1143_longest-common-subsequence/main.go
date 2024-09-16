package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for j := len(dp) - 2; j >= 0; j-- {
		for k := len(dp[j]) - 2; k >= 0; k-- {
			if text1[j] == text2[k] {
				dp[j][k] = dp[j+1][k+1] + 1
			} else {
				if dp[j+1][k] > dp[j][k+1] {
					dp[j][k] = dp[j+1][k]
				} else {
					dp[j][k] = dp[j][k+1]
				}
			}
		}
	}

	return dp[0][0]
}
