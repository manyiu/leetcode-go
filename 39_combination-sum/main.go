package combinationsum

func backtrack(candidates []int, target int, step int, path []int, res *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := step; i < len(candidates); i++ {
		backtrack(candidates, target-candidates[i], i, append(path, candidates[i]), res)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	backtrack(candidates, target, 0, []int{}, &res)

	return res
}
