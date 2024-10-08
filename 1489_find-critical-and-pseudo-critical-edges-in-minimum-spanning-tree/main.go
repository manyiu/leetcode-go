package findcriticalandpseudocriticaledgesinminimumspanningtree

import (
	"container/heap"
)

type minHeap [][]int

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
	*mh = append(*mh, x.([]int))
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

func (uf unionFind) Find(node int) int {
	curr := node

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
		uf.parent[rootA] = b
	} else if uf.rank[a] > uf.rank[rootB] {
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

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	mstWeight := 0

	{
		currEdge := make([][]int, len(edges))
		copy(currEdge, edges)
		mh := minHeap(currEdge)
		heap.Init(&mh)

		uf := newUnionFind(n)

		for len(mh) > 0 {
			popped := heap.Pop(&mh).([]int)
			a, b, weight := popped[0], popped[1], popped[2]

			if uf.Union(a, b) {
				mstWeight += weight
			}
		}
	}

	criticalEdges := []int{}
	pseudoCriticalEdges := []int{}

	{
		for currIndex, currEdge := range edges {
			{
				edgesWithoutCurrEdge := [][]int{}

				for i, edge := range edges {
					if i != currIndex {
						edgesWithoutCurrEdge = append(edgesWithoutCurrEdge, edge)
					}
				}

				mh := minHeap(edgesWithoutCurrEdge)
				heap.Init(&mh)

				uf := newUnionFind(n)

				totalWeight := 0

				for len(mh) > 0 {
					popped := heap.Pop(&mh).([]int)
					a, b, weight := popped[0], popped[1], popped[2]

					if uf.Union(a, b) {
						totalWeight += weight
					}
				}

				allConnected := true

				for i := 1; i < n; i++ {
					if uf.Find(i-1) != uf.Find(i) {
						allConnected = false
					}
				}

				if totalWeight > mstWeight || !allConnected {
					criticalEdges = append(criticalEdges, currIndex)
					continue
				}
			}

			{
				edgesWithoutCurrEdge := [][]int{}

				for i, edge := range edges {
					if i != currIndex {
						edgesWithoutCurrEdge = append(edgesWithoutCurrEdge, edge)
					}
				}

				mh := minHeap(edgesWithoutCurrEdge)
				heap.Init(&mh)

				uf := newUnionFind(n)
				uf.Union(currEdge[0], currEdge[1])
				totalWeight := currEdge[2]

				for len(mh) > 0 {
					popped := heap.Pop(&mh).([]int)
					a, b, weight := popped[0], popped[1], popped[2]

					if uf.Union(a, b) {
						totalWeight += weight
					}
				}

				if totalWeight == mstWeight {
					pseudoCriticalEdges = append(pseudoCriticalEdges, currIndex)
				}
			}
		}
	}

	return [][]int{criticalEdges, pseudoCriticalEdges}
}
