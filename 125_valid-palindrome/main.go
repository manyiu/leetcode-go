package validpalindrome

func isPalindrome(s string) bool {
	l := []rune{}

	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			l = append(l, v)
		} else if v >= 'A' && v <= 'Z' {
			l = append(l, v+('a'-'A'))
		}
	}

	if len(l) <= 1 {
		return true
	}

	for i := 0; i <= (len(l) / 2); i++ {
		if l[i] != l[len(l)-1-i] {
			return false
		}
	}

	return true
}
