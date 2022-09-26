package longestpalindrome

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	l := 0

	for _, c := range s {
		if m[c] == 1 {
			m[c]--
			l += 2
		} else {
			m[c] = 1
		}
	}

	if l < len(s) {
		l++
	}

	return l
}
