package palindromicsubstrings

func countSubstrings(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		for l, r := i, i; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				result++
			} else {
				break
			}
		}

		for l, r := i, i+1; l >= 0 && r < len(s); l, r = l-1, r+1 {
			if s[l] == s[r] {
				result++
			} else {
				break
			}
		}
	}

	return result
}
