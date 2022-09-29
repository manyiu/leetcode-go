package majorityelement

func majorityElement(nums []int) int {
	c, n := 0, 0

	for _, v := range nums {
		if c == 0 {
			n = v
		}

		if v == n {
			c++
		} else {
			c--
		}
	}

	return n
}
