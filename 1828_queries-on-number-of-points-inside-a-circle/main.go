package main

import (
	"math"
)

func countPoints(points [][]int, queries [][]int) []int {
	r := make([]int, len(queries))

	for i, q := range queries {
		for _, p := range points {
			if math.Pow(float64(p[0]-q[0]), 2)+math.Pow(float64(p[1]-q[1]), 2) <= math.Pow(float64(q[2]), 2) {
				r[i]++
			}
		}
	}

	return r
}

func main() {}
