package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)

	for i := 0; i < len(text1); i++ {
		row := make([]int, len(text2)+1)

		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				row[j+1] = dp[j] + 1
			} else {
				row[j+1] = max(dp[j+1], row[j])
			}
		}

		dp = row
	}

	return dp[len(dp)-1]
}
