package trappingrainwater

func trap(height []int) int {
	total := 0

	for {
		l, r := 0, len(height)-1

		for l < r {
			if height[l] != 0 && height[r] != 0 {
				break
			}

			if height[l] == 0 {
				l++
			}

			if height[r] == 0 {
				r--
			}
		}

		if r-l <= 1 {
			break
		}

		count := 0

		for i := l; i <= r; i++ {
			if height[i] == 0 {
				count++
			}
		}

		total += count

		for j := 0; j < len(height); j++ {
			if height[j] > 0 {
				height[j] -= 1
			}
		}
	}

	return total
}
