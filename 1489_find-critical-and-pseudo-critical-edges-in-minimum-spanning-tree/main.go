package findcriticalandpseudocriticaledgesinminimumspanningtree

import (
	"container/heap"
)

type minHeap [][2]int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.([2]int))
}

func (mh *minHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	adjs := make([][][2]int, n)

	mstWeight := 0

	{
		for _, edge := range edges {
			adjs[edge[0]] = append(adjs[edge[0]], [2]int{edge[1], edge[2]})
			adjs[edge[1]] = append(adjs[edge[1]], [2]int{edge[0], edge[2]})
		}

		mh := &minHeap{
			{0, 0},
		}

		visit := make([]bool, n)
		edgeBuildCount := -1

		for edgeBuildCount < n-1 {
			popped := heap.Pop(mh).([2]int)
			node, weight := popped[0], popped[1]

			if visit[node] {
				continue
			}

			visit[node] = true
			mstWeight += weight
			edgeBuildCount++

			for _, adj := range adjs[node] {
				nextNode, nextWeight := adj[0], adj[1]
				if !visit[nextNode] {
					heap.Push(mh, [2]int{nextNode, nextWeight})
				}
			}
		}
	}

	criticalEdge := []int{}
	pseudoCriticalEdge := []int{}

	for currEdgeIndex, currEdge := range edges {
		adjWithoutCurr := make([][][2]int, n)

		{

			for j, edge := range edges {
				if currEdgeIndex != j {
					adjWithoutCurr[edge[0]] = append(adjWithoutCurr[edge[0]], [2]int{edge[1], edge[2]})
					adjWithoutCurr[edge[1]] = append(adjWithoutCurr[edge[1]], [2]int{edge[0], edge[2]})
				}
			}

			startNode := 0

			if currEdgeIndex == 0 {
				startNode = 1

			}
			mh := &minHeap{
				{startNode, 0},
			}

			visit := make([]bool, n)
			edgeBuildCount := -1
			totalWeight := 0

			for edgeBuildCount < n-1 && len(*mh) > 0 {
				popped := heap.Pop(mh).([2]int)

				node, weight := popped[0], popped[1]

				if visit[node] {
					continue
				}

				visit[node] = true
				totalWeight += weight
				edgeBuildCount++

				for _, adj := range adjWithoutCurr[node] {
					nextNode, nextWeight := adj[0], adj[1]
					if !visit[nextNode] {
						heap.Push(mh, [2]int{nextNode, nextWeight})
					}
				}
			}

			visitCount := 0

			for _, v := range visit {
				if v {
					visitCount++
				}
			}

			if totalWeight > mstWeight || visitCount != n {
				criticalEdge = append(criticalEdge, currEdgeIndex)
				continue
			}
		}

		{
			mh := &minHeap{}
			visit := make([]bool, n)
			edgeBuildCount := 1
			totalWeight := currEdge[2]
			visit[currEdge[0]] = true
			visit[currEdge[1]] = true

			for _, adj := range adjWithoutCurr[currEdge[0]] {
				heap.Push(mh, adj)
			}
			for _, adj := range adjWithoutCurr[currEdge[1]] {
				heap.Push(mh, adj)
			}

			for edgeBuildCount < n-1 && len(*mh) > 0 {
				popped := heap.Pop(mh).([2]int)
				node, weight := popped[0], popped[1]

				if visit[node] {
					continue
				}

				visit[node] = true
				edgeBuildCount++
				totalWeight += weight

				for _, adj := range adjWithoutCurr[node] {
					nextNode, nextWeight := adj[0], adj[1]
					if !visit[nextNode] {
						heap.Push(mh, [2]int{nextNode, nextWeight})
					}
				}

			}

			if totalWeight == mstWeight {
				pseudoCriticalEdge = append(pseudoCriticalEdge, currEdgeIndex)
			}
		}

	}

	return [][]int{criticalEdge, pseudoCriticalEdge}
}
