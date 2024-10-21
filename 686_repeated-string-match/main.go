package repeatedstringmatch

func isSubString(a, b string) bool {
	if len(b) > len(a) {
		return false
	}

	for i := 0; i <= len(a)-len(b); i++ {
		if a[i:i+len(b)] == b {
			return true
		}
	}

	return false
}

func repeatedStringMatch(a string, b string) int {
	maxTries := len(b)/len(a) + 2
	count := 1
	repeatedA := a

	for count <= maxTries {
		if isSubString(repeatedA, b) {
			return count
		}

		repeatedA += a
		count++
	}

	return -1
}
