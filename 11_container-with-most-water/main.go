package containerwithmostwater

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	m := 0

	for l < r {
		area := (r - l) * min(height[l], height[r])
		m = max(m, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return m
}
