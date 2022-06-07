package validanagram

import "strings"

func isAnagram(s string, t string) bool {
	for _, c := range s {
		b := len(t)

		t = strings.Replace(t, string(c), "", 1)

		if len(t) == b {
			return false
		}
	}

	return len(t) == 0
}
