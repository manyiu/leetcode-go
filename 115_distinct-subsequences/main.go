package distinctsubsequences

func dp(s, t string, i, j int, cache [][]int) int {
	if len(t) == j {
		return 1
	}
	if len(s) == i {
		return 0
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	if s[i] == t[j] {
		cache[i][j] = dp(s, t, i+1, j+1, cache) + dp(s, t, i+1, j, cache)
	} else {
		cache[i][j] = dp(s, t, i+1, j, cache)
	}

	return cache[i][j]
}

func numDistinct(s string, t string) int {
	cache := make([][]int, len(s)+1)

	for i := 0; i < len(cache); i++ {
		row := make([]int, len(t)+1)

		for j := 0; j < len(row); j++ {
			row[j] = -1
		}

		cache[i] = row
	}

	return dp(s, t, 0, 0, cache)
}
