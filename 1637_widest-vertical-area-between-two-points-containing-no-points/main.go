package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
    xs := make([]int, len(points))

	for i, point := range points {
		xs[i] = point[0]
	}

	sort.Ints(xs)

	r := 0

	for i := 0; i < len(xs) - 1; i++ {
		d := xs[i+1] - xs[i]
		if d > r {
			r = d
		}
	}

	return r
}

func main() {}
