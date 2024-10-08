package mincosttoconnectallpoints

import (
	"container/heap"
	"math"
)

type minHeap [][3]int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i][2] < mh[j][2]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.([3]int))
}

func (mh *minHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

type unionFind struct {
	parent []int
	rank   []int
}

func (uf unionFind) Find(val int) int {
	curr := val

	for curr != uf.parent[curr] {
		uf.parent[curr] = uf.parent[uf.parent[curr]]
		curr = uf.parent[curr]
	}

	return curr
}

func (uf unionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}

	return true
}

func newUnionFind(n int) unionFind {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return unionFind{
		parent,
		rank,
	}
}

func distance(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func minCostConnectPoints(points [][]int) int {
	edges := minHeap{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, [3]int{i, j, distance(points[i], points[j])})
		}
	}

	heap.Init(&edges)

	uf := newUnionFind(len(points))

	edgeCount := 0
	totalCost := 0

	for edgeCount < len(points)-1 {
		popped := heap.Pop(&edges).([3]int)
		a, b, cost := popped[0], popped[1], popped[2]

		if uf.Union(a, b) {
			totalCost += cost
			edgeCount++
		}
	}

	return totalCost
}
