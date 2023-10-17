package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	r := make([]int, len(nums))
	r[0] = 1

	// nums : []int{a, b, c, d, e}
	//    r : []int{1, 0, 0, 0, 0}

	for i := 1; i < len(nums); i++ {
		r[i] = nums[i-1] * r[i-1]
	}

	// r : []int{   1,    a,   ab,  abc, abcd} <= from above
	// n : []int{bcde,  cde,   de,    e,    1} <= need this to complete
	// w : []int{bcde, acde, abde, abce, abcd} <= wanted

	s := 1

	for i := len(nums) - 2; i >= 0; i-- {
		s *= nums[i+1]
		r[i] *= s
	}

	return r
}
