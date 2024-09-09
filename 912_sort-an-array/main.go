package sortanarray

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))

	for i, j, k := 0, 0, 0; i < len(a) || j < len(b); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] > b[j] {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}
	}

	return c
}

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	a := sortArray(nums[:len(nums)/2])
	b := sortArray(nums[len(nums)/2:])

	return merge(a, b)
}
