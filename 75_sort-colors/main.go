package sortcolors

func sortColors(nums []int) {
	c := []int{
		0,
		0,
		0,
	}

	for _, num := range nums {
		c[num]++
	}

	n := 0

	for i, v := range c {
		for j := 0; j < v; j++ {
			nums[n] = i
			n++
		}
	}
}
