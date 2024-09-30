package queuereconstructionbyheight

import (
	"sort"
)

type SegmentTree struct {
	start, end, sum int
	left, right     *SegmentTree
}

func NewSegmentTree(nums []int, start, end int) *SegmentTree {
	if start == end {
		return &SegmentTree{
			start: start,
			end:   end,
			sum:   nums[start],
		}
	}

	m := (start + end) / 2

	left := NewSegmentTree(nums, start, m)
	right := NewSegmentTree(nums, m+1, end)

	sum := left.sum + right.sum

	return &SegmentTree{
		start,
		end,
		sum,
		left,
		right,
	}
}

func (st *SegmentTree) Update(index, value int) {
	if st.start == index && st.end == index {
		st.sum = value
		return
	}

	m := (st.start + st.end) / 2

	if index <= m {
		st.left.Update(index, value)
	} else {
		st.right.Update(index, value)
	}

	st.sum = st.left.sum + st.right.sum
}

func (st *SegmentTree) Query(start, end int) int {
	if st.start == start && st.end == end {
		return st.sum
	}

	m := (st.start + st.end) / 2

	if end <= m {
		return st.left.Query(start, end)
	}

	if start > m {
		return st.right.Query(start, end)
	}

	return st.left.Query(start, m) + st.right.Query(m+1, end)
}

func (st *SegmentTree) Find(k int) int {
	if st.start == st.end {
		return st.start
	}

	if st.left.sum >= k {
		return st.left.Find(k)
	} else {
		return st.right.Find(k - st.left.sum)
	}
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			return people[a][1] > people[b][1]
		}

		return people[a][0] < people[b][0]
	})

	v := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		v[i] = 1
	}

	st := NewSegmentTree(v, 0, len(people)-1)

	result := make([][]int, len(people))

	for j := 0; j < len(people); j++ {
		pos := st.Find(people[j][1] + 1)
		result[pos] = people[j]
		st.Update(pos, 0)
	}

	return result
}
