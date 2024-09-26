package accountsmerge

import "sort"

type UnionFind struct {
	parent []int
	rank   []int
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

func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind(len(accounts))

	emailToIndex := map[string]int{}

	for i, account := range accounts {
		for _, email := range account[1:] {
			if existedIndex, ok := emailToIndex[email]; ok {
				uf.Union(i, existedIndex)
			} else {
				emailToIndex[email] = i
			}
		}
	}

	merged := map[int][]string{}

	for email, accountIndex := range emailToIndex {
		root := uf.Find(accountIndex)
		merged[root] = append(merged[root], email)
	}

	result := [][]string{}

	for accountIndex, emails := range merged {
		sort.Strings(emails)
		result = append(result, append([]string{accounts[accountIndex][0]}, emails...))
	}

	return result
}
