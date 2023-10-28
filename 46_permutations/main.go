package permutations

func recursive(allocated []int, available []int, result *[][]int) {
	if len(available) == 0 {
		*result = append(*result, allocated)
		return
	}

	for i, v := range available {
		recursive(append(allocated, v), append(append([]int{}, available[:i]...), available[i+1:]...), result)
	}
}

func permute(nums []int) [][]int {
	r := [][]int{}
	recursive([]int{}, nums, &r)

	return r
}
