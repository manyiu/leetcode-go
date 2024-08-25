package combinationsum

func dfs(candidates []int, target int, index int, member []int, accu *[][]int) {
	if target < 0 || index >= len(candidates) {
		return
	}

	if target == 0 {
		*accu = append(*accu, append([]int{}, member...))
		return
	}

	dfs(candidates, target-candidates[index], index, append(member, candidates[index]), accu)
	dfs(candidates, target, index+1, member, accu)
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	dfs(candidates, target, 0, []int{}, &result)

	return result
}
