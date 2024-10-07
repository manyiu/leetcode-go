package mincosttoconnectallpoints

import (
	"container/heap"
	"math"
)

type MinHeap [][3]int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][2] < mh[j][2]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([3]int))
}

func (mh *MinHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {
	visit := make([]bool, len(points))

	minHeap := &MinHeap{
		// { {pointA, pointB, distance} }
		{0, 0, 0},
	}

	totalCost := 0

	for len(*minHeap) > 0 {
		popped := heap.Pop(minHeap).([3]int)

		poppedPoint, poppedCost := popped[1], popped[2]

		if visit[poppedPoint] {
			continue
		}

		visit[poppedPoint] = true

		totalCost += poppedCost

		x0, y0 := points[poppedPoint][0], points[poppedPoint][1]

		for i, point := range points {
			if !visit[i] {
				x, y := point[0], point[1]
				heap.Push(minHeap, [3]int{poppedPoint, i, int(math.Abs(float64(x-x0)) + math.Abs(float64(y-y0)))})
			}
		}
	}

	return totalCost
}
