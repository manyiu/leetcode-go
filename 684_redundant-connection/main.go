package redundantconnection

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n+1)
	rank := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	return &UnionFind{
		parent,
		rank,
	}
}

func (uf *UnionFind) Find(node int) int {
	curr := node

	for curr != uf.parent[curr] {
		uf.parent[curr] = uf.parent[uf.parent[curr]]
		curr = uf.parent[curr]
	}

	return curr
}

func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}

	return true
}

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges))

	result := []int{}

	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			result = edge
		}
	}

	return result
}
