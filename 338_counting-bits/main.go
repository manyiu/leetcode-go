package countingbits

func countBits(n int) []int {
	output := make([]int, n+1)

	for i := 0; i <= n; i++ {
		c := 0

		for j := i; j > 0; j = j >> 1 {
			if j&1 == 1 {
				c++
			}
		}

		output[i] = c
	}

	return output
}
