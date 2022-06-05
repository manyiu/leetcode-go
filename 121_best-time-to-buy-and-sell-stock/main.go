package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	p := 0

	for _, v := range prices {
		if v < min {
			min = v
		}

		if v-min > p {
			p = v - min
		}
	}

	return p
}
