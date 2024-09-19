package validpalindrome

func validChar(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}

	return false
}

func encodeChar(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	} else {
		return c
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for !validChar(s[l]) {
			l++
		}

		for !validChar(s[r]) {
			r--
		}

		if encodeChar(s[l]) != encodeChar(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}
