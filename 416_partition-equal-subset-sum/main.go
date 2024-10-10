package partitionequalsubsetsum

import "strconv"

func dfs(nums []int, i, target int, cache map[string]bool) bool {
	if i >= len(nums) || target < 0 {
		return false
	}

	if target-nums[i] == 0 {
		return true
	}

	key := strconv.Itoa(i) + "-" + strconv.Itoa(target)

	cache[key] = dfs(nums, i+1, target, cache) || dfs(nums, i+1, target-nums[i], cache)

	return cache[key]
}

func canPartition(nums []int) bool {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	cache := map[string]bool{}

	return dfs(nums, 0, target, cache)
}
