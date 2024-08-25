package subsets

func dfs(nums []int, index int, subset []int, accu *[][]int) {
	if index >= len(nums) {
		*accu = append(*accu, subset)
		return
	}

	dfs(nums, index+1, subset, accu)

	dfs(nums, index+1, append(subset, nums[index]), accu)
}

func subsets(nums []int) [][]int {
	result := [][]int{}

	dfs(nums, 0, []int{}, &result)

	return result
}
