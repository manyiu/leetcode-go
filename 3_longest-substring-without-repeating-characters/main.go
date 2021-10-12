package main

func lengthOfLongestSubstring(s string) int {
	m := 0
	b := 0
	h := map[rune]int{}

	for i, v := range s {
		if p, ok := h[v]; ok && p >= b {
			b = p + 1
		} else if i-b+1 > m {
			m = i - b + 1
		}

		h[v] = i
	}

	return m
}

func main() {}
