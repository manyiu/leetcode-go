package permutationsii

import (
	"slices"
)

func backtrack(available []int, curr []int, result *[][]int) {
	if len(available) == 0 {
		*result = append(*result, curr)
		return
	}

	for i := 0; i < len(available); i++ {
		if i >= 1 && available[i] == available[i-1] {
			continue
		}

		copyCurr := make([]int, len(curr))
		copy(copyCurr, curr)
		availableCopy := make([]int, len(available))
		copy(availableCopy, available)
		nextAvailable := append(availableCopy[:i], availableCopy[i+1:]...)
		backtrack(nextAvailable, append(copyCurr, available[i]), result)
	}
}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)

	result := [][]int{}

	backtrack(nums, []int{}, &result)

	return result
}
