package combinationsum

func backtrack(candidates, set []int, target, i int, result *[][]int) {
	if target == 0 {
		*result = append(*result, set)
		return
	}

	if target < 0 {
		return
	}

	for j := i; j < len(candidates); j++ {
		backtrack(candidates, append([]int{}, append(set, candidates[j])...), target-candidates[j], j, result)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	backtrack(candidates, []int{}, target, 0, &result)

	return result
}
