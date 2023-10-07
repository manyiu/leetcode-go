package sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	a := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		t, l, r := -nums[i], i+1, len(nums)-1

		for l < r {
			s := nums[l] + nums[r]

			if s == t {
				a = append(a, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s < t {
				l++
			} else {
				r--
			}
		}
	}

	return a
}
