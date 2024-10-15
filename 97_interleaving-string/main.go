package interleavingstring

func backtrack(s1, s2, s3 string, i1, i2, i3 int) bool {
	if i3 >= len(s3) && (i1 < len(s1) || i2 < len(s2)) {
		return false
	}

	if i3 >= len(s3) {
		return true
	}

	if i1 < len(s1) && s1[i1] == s3[i3] && i2 < len(s2) && s2[i2] == s3[i3] {
		return backtrack(s1, s2, s3, i1+1, i2, i3+1) || backtrack(s1, s2, s3, i1, i2+1, i3+1)
	}

	if i1 < len(s1) && s1[i1] == s3[i3] {
		return backtrack(s1, s2, s3, i1+1, i2, i3+1)
	}

	if i2 < len(s2) && s2[i2] == s3[i3] {
		return backtrack(s1, s2, s3, i1, i2+1, i3+1)
	}

	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return backtrack(s1, s2, s3, 0, 0, 0)
}
