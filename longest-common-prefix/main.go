package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	v := ""

	if len(strs) == 0 {
		return v
	}

	min := len(strs[0])

	if min == 0 {
		return v
	}

	cn := true

	for _, s := range strs {
		if len(s) < min {
			min = len(s)
		}
	}

	for i := 0; i < min; i++ {
		c := strs[0][i]
		for _, w := range strs {
			if w[i] != c {
				cn = false
				break
			}
		}
		if cn == false {
			break
		}
		v = v + string(c)
	}

	return v
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
