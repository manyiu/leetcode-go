package longestconsecutivesequence

type UnionFind struct {
	parent, rank []int
}

func (uf *UnionFind) Find(node int) int {
	curr := node

	for curr != uf.parent[curr] {
		uf.parent[curr] = uf.parent[uf.parent[curr]]
		curr = uf.parent[curr]
	}

	return curr
}

func (uf *UnionFind) Union(a, b int) {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return
	}

	if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UnionFind{
		parent,
		rank,
	}
}

func longestConsecutive(nums []int) int {
	uf := NewUnionFind(len(nums))

	valueToIndexMap := map[int]int{}

	for i, num := range nums {
		valueToIndexMap[num] = i

		prevIndex, hasPrev := valueToIndexMap[num-1]
		nextIndex, hasNext := valueToIndexMap[num+1]

		if hasPrev {
			uf.Union(i, prevIndex)
		}

		if hasNext {
			uf.Union(i, nextIndex)
		}
	}

	groupedCount := map[int]int{}
	max := 0

	for _, index := range valueToIndexMap {
		root := uf.Find(index)

		groupedCount[root]++

		if groupedCount[root] > max {
			max = groupedCount[root]
		}
	}

	return max
}
