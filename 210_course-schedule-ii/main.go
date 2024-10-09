package coursescheduleii

func dfs(adj [][]int, src int, visit, path []bool, topSort *[]int) bool {
	if path[src] {
		return false
	}

	if visit[src] {
		return true
	}

	path[src] = true
	visit[src] = true

	for _, pre := range adj[src] {
		if !dfs(adj, pre, visit, path, topSort) {
			return false
		}
	}

	*topSort = append(*topSort, src)
	path[src] = false

	return true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)

	for _, pre := range prerequisites {
		adj[pre[0]] = append(adj[pre[0]], pre[1])
	}

	visit := make([]bool, numCourses)
	path := make([]bool, numCourses)
	topSort := []int{}

	for i := 0; i < numCourses; i++ {
		if !dfs(adj, i, visit, path, &topSort) {
			return []int{}
		}
	}

	return topSort
}
