package networkdelaytime

import (
	"container/heap"
)

type MinHeap [][2]int

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([2]int))
}

func (mh *MinHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {

	adj := map[int][][2]int{
		// src: {dest, dist}
	}

	for _, time := range times {
		adj[time[0]] = append(adj[time[0]], [2]int{time[1], time[2]})
	}

	minHeap := &MinHeap{
		// {dest, dist}
		{k, 0},
	}

	shortest := map[int]int{
		// dest: dist
	}

	count := 0

	for len(*minHeap) > 0 {
		popped := heap.Pop(minHeap).([2]int)
		poppedDest, poppedDist := popped[0], popped[1]

		if _, ok := shortest[poppedDest]; ok {
			continue
		}

		shortest[poppedDest] = poppedDist
		count++

		if count == n {
			return poppedDist
		}

		routes, hasRoute := adj[poppedDest]

		if !hasRoute {
			continue
		}

		for _, route := range routes {
			routeDest, routeDist := route[0], route[1]

			_, hasShortest := shortest[routeDest]

			if !hasShortest {
				heap.Push(minHeap, [2]int{routeDest, routeDist + poppedDist})
			}
		}
	}

	return -1
}
