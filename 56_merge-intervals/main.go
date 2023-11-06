package mergeintervals

import (
	"sort"
)

func sortIntervals(i [][]int) {
	sort.Slice(i, func(a, b int) bool {
		return i[a][0] < i[b][0]
	})
}

func merge(intervals [][]int) [][]int {
	sortIntervals(intervals)

	m := [][]int{intervals[0]}

	for _, interval := range intervals {
		c := m[len(m)-1]

		if interval[0] <= c[1] {
			if interval[1] > c[1] {
				c[1] = interval[1]
			}
		} else {
			m = append(m, interval)
		}
	}

	return m
}
