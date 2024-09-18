package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	l, r := 0, 0

	m := map[int]bool{
		nums[0]: true,
	}

	for r < len(nums)-1 {
		r++

		if r-l > k {
			delete(m, nums[l])
			l++
		}

		if _, existed := m[nums[r]]; existed {
			return true
		}

		m[nums[r]] = true
	}

	return false
}
