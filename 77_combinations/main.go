package combinations

func backtrack(n, k, i int, set []int, result *[][]int) {
	if len(set) >= k {
		*result = append(*result, set)
		return
	}

	if i > n {
		return
	}

	for j := i; j <= n; j++ {
		backtrack(n, k, j+1, append([]int{}, append(set, j)...), result)
	}
}

func combine(n int, k int) [][]int {
	result := [][]int{}

	backtrack(n, k, 1, []int{}, &result)

	return result
}
