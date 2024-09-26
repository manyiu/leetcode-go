package numberofconnectedcomponentsinanundirectedgraph

type UnionFind struct {
	parent []int
	rank   []int
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
		uf.parent[rootA]++
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

func countComponents(n int, edges [][]int) int {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	count := 0
	hash := map[int]bool{}

	for i := 0; i < n; i++ {
		root := uf.Find(i)
		if _, ok := hash[root]; !ok {
			hash[root] = true
			count++
		}
	}

	return count
}
