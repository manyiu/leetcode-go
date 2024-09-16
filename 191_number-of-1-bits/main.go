package numberof1bits

func hammingWeight(n int) int {
	c := 0

	for n > 0 {
		if n&1 == 1 {
			c++
		}

		n = n >> 1
	}

	return c
}
