package subsetsii

import "sort"

func backtrack(nums []int, i int, subset []int, result *[][]int) {
	if i >= len(nums) {
		*result = append(*result, append([]int{}, subset...))
		return
	}

	backtrack(nums, i+1, append(subset, nums[i]), result)

	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}

	backtrack(nums, i+1, subset, result)
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	backtrack(nums, 0, []int{}, &result)

	return result
}
