package main

func search(nums []int, target int) int {
	l := 0
	u := len(nums) - 1

	for l <= u {
		m := (l + u) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			u = m - 1
		}
	}

	return -1
}

func main() {}
