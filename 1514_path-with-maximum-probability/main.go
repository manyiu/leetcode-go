package pathwithmaximumprobability

import (
	"container/heap"
)

type Path struct {
	node int
	prob float64
}

type MaxHeap []Path

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i].prob > mh[j].prob
}

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(Path))
}

func (mh *MaxHeap) Pop() interface{} {
	n := len(*mh)
	x := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	adj := map[int][]Path{}

	for i, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], Path{
			node: edge[1], prob: succProb[i],
		})
		adj[edge[1]] = append(adj[edge[1]], Path{
			node: edge[0], prob: succProb[i],
		})
	}

	maxHeap := &MaxHeap{
		Path{node: start_node, prob: 1},
	}

	visited := make([]bool, n)

	for len(*maxHeap) > 0 {
		popped := heap.Pop(maxHeap).(Path)
		node, prob := popped.node, popped.prob

		if node == end_node {
			return prob
		}

		visited[node] = true

		for _, path := range adj[node] {
			if !visited[path.node] {
				heap.Push(maxHeap, Path{node: path.node, prob: path.prob * prob})
			}
		}
	}

	return 0
}
