package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	output := 0
	l, r := 0, k-1

	for r < len(arr) {
		sum := 0

		for i := l; i <= r; i++ {
			sum += arr[i]
		}

		if sum >= k*threshold {
			output += 1
		}

		l++
		r++
	}

	return output
}
