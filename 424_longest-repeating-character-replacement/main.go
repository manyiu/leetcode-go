package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	l, r := 0, 0
	maxLength := 0
	maxCount := 0
	count := make([]int, 26)

	for r < len(s) {
		count[s[r]-'A']++

		maxCount = max(maxCount, count[s[r]-'A'])

		if r-l+1-maxCount <= k {
			maxLength = max(maxLength, r-l+1)
		} else {
			for r-l+1-maxCount > k {
				count[s[l]-'A']--
				l++
			}
		}

		r++
	}

	return maxLength
}
