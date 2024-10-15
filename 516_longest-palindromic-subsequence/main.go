package longestpalindromicsubsequence

func helper(s string, l, r int, cache map[[2]int]int) int {
	if v, ok := cache[[2]int{l, r}]; ok {
		return v
	}

	if l < 0 || r >= len(s) {
		return 0
	}

	if s[l] == s[r] {
		length := 1

		if l != r {
			length = 2
		}

		cache[[2]int{l, r}] = length + helper(s, l-1, r+1, cache)
	} else {
		cache[[2]int{l, r}] = max(helper(s, l-1, r, cache), helper(s, l, r+1, cache))
	}

	return cache[[2]int{l, r}]
}

func longestPalindromeSubseq(s string) int {
	output := 0
	cache := map[[2]int]int{}

	for i := 0; i < len(s); i++ {
		output = max(output, helper(s, i, i, cache), helper(s, i, i+1, cache))
	}

	return output
}
