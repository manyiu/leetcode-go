package courseschedule

func dfs(src int, adj [][]int, visit, path []bool, topSort *[]int) bool {
	if path[src] {
		return false
	}

	if visit[src] {
		return true
	}

	visit[src] = true
	path[src] = true

	for _, i := range adj[src] {
		if !dfs(i, adj, visit, path, topSort) {
			return false
		}
	}

	*topSort = append(*topSort, src)

	path[src] = false

	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)

	for _, pre := range prerequisites {
		adj[pre[0]] = append(adj[pre[0]], pre[1])
	}

	visit := make([]bool, numCourses)
	path := make([]bool, numCourses)
	topSort := []int{}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, adj, visit, path, &topSort) {
			return false
		}
	}

	return true
}
