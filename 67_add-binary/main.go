package addbinary

func addBinary(a string, b string) string {
	answer := ""
	carry := 0

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(a[i] - '0')
		}

		if j >= 0 {
			carry += int(b[j] - '0')
		}

		answer = string(rune(carry%2+'0')) + answer
		carry /= 2
	}

	if carry > 0 {
		answer = "1" + answer
	}

	return answer
}
