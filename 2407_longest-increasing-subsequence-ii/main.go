package longestincreasingsubsequenceii

type SegmentTree struct {
	start, end, max int
	left, right     *SegmentTree
}

func NewSegmentTree(nums []int, start, end int) *SegmentTree {
	if start == end {
		return &SegmentTree{
			start: start,
			end:   end,
			max:   nums[start],
		}
	}

	m := (start + end) / 2

	left := NewSegmentTree(nums, start, m)
	right := NewSegmentTree(nums, m+1, end)

	max := max(left.max, right.max)

	return &SegmentTree{
		start: start,
		end:   end,
		left:  left,
		right: right,
		max:   max,
	}
}

func (st *SegmentTree) Update(index, value int) {

	if st.start == st.end && st.start == index {
		st.max = value

		return
	}

	m := (st.start + st.end) / 2

	if index <= m {
		st.left.Update(index, value)
	} else {
		st.right.Update(index, value)
	}

	st.max = max(st.left.max, st.right.max)
}

func (st *SegmentTree) MaxRange(start, end int) int {
	if st.start == start && st.end == end {
		return st.max
	}

	m := (st.start + st.end) / 2

	if end <= m {
		return st.left.MaxRange(start, end)
	}

	if start > m {
		return st.right.MaxRange(start, end)
	}

	return max(st.left.MaxRange(start, m), st.right.MaxRange(m+1, end))
}

func lengthOfLIS(nums []int, k int) int {
	max := 0

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	l := make([]int, max)

	st := NewSegmentTree(l, 0, max-1)

	for _, num := range nums {
		startIndex := num - k - 1

		if startIndex < 0 {
			startIndex = 0
		}

		endIndex := num - 1 - 1

		if endIndex < 0 {
			st.Update(0, 1)
		} else {
			subMax := st.MaxRange(startIndex, endIndex)
			st.Update(num-1, subMax+1)
		}

	}

	result := st.MaxRange(0, max-1)

	return result
}
