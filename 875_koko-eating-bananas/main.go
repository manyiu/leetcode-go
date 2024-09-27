package kokoeatingbananas

func getHourToEat(piles []int, eatingSpeed int) int {
	hour := 0

	for _, pile := range piles {
		count := pile / eatingSpeed

		if pile%eatingSpeed > 0 {
			count++
		}

		hour += count
	}

	return hour
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1

	for i := 0; i < len(piles); i++ {
		if piles[i] > r {
			r = piles[i]
		}
	}

	for l <= r {

		m := (l + r) / 2

		hourToEat := getHourToEat(piles, m)

		if hourToEat <= h {
			r = m - 1
		} else {
			l = m + 1
		}

	}

	return l
}
