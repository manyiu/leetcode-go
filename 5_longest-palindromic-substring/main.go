package longestpalindromicsubstring

func helper(s string, l, r int) [3]int {
	ll, lr, length := -1, -1, 0

	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			length = r - l + 1
			ll = l
			lr = r
			l--
			r++
		} else {
			break
		}
	}

	return [3]int{length, ll, lr}
}

func longestPalindrome(s string) string {
	longest := 0
	l, r := -1, -1

	for i := 0; i < len(s); i++ {
		odd := helper(s, i, i)
		even := helper(s, i, i+1)

		if odd[0] > longest {
			longest = odd[0]
			l = odd[1]
			r = odd[2]
		}

		if even[0] > longest {
			longest = even[0]
			l = even[1]
			r = even[2]
		}
	}

	return s[l : r+1]
}
