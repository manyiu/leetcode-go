package trappingrainwater

func trap(height []int) int {
	l, r := 0, len(height)-1
	maxLeft, maxRight := height[l], height[r]
	total := 0

	for l < r {
		if maxLeft < maxRight {
			l++
			maxLeft = max(maxLeft, height[l])
			total += maxLeft - height[l]
		} else {
			r--
			maxRight = max(maxRight, height[r])
			total += maxRight - height[r]
		}
	}

	return total
}
