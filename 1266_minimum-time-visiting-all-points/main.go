package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	d := 0

	for i := 0; i < len(points)-1; i++ {
		dx := points[i][0] - points[i+1][0]
		dy := points[i][1] - points[i+1][1]

		d += int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))
	}

	return d
}

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
}
