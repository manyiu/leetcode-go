package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	longest := 0
	m := map[byte]int{}

	for r < len(s) {
		if i, existed := m[s[r]]; existed && i >= l {
			l = i + 1
		}

		m[s[r]] = r

		if r-l+1 > longest {
			longest = r - l + 1
		}

		r++
	}

	return longest
}
