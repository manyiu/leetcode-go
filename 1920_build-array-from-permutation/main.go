package main

func buildArray(nums []int) []int {
	p := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		p[i] = nums[nums[i]]
	}

	return p
}

func main() {}
