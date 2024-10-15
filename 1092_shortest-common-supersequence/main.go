package shortestcommonsupersequence

func shortestCommonSupersequence(str1 string, str2 string) string {
	dp := make([][]int, len(str1)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(str2)+1)
	}

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	output := ""

	i, j := len(str1)-1, len(str2)-1

	for i >= 0 && j >= 0 {
		if str1[i] == str2[j] {
			output = string(str1[i]) + output
			i--
			j--

			continue
		}

		if dp[i+1][j] > dp[i][j+1] {
			output = string(str2[j]) + output
			j--
		} else {
			output = string(str1[i]) + output
			i--
		}
	}

	for i >= 0 {
		output = string(str1[i]) + output
		i--
	}

	for j >= 0 {
		output = string(str2[j]) + output
		j--
	}

	return output
}
