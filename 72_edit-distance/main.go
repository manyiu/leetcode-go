package editdistance

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}

	for i := 0; i < len(word1); i++ {
		row := make([]int, len(word2)+1)
		row[0] = i + 1

		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				row[j+1] = dp[j]
			} else {
				row[j+1] = min(dp[j], dp[j+1], row[j]) + 1
			}
		}

		dp = row
	}

	return dp[len(dp)-1]
}
