package coursescheduleiv

func dfs(adj [][]int, course int, allPrerequisites []map[int]bool, visit []bool) map[int]bool {
	if allPrerequisites[course] == nil {
		allPrerequisites[course] = map[int]bool{}

		for _, pre := range adj[course] {
			subMap := dfs(adj, pre, allPrerequisites, visit)

			for k, v := range subMap {
				if v {
					allPrerequisites[course][k] = true
				}
			}

			allPrerequisites[course][pre] = true
		}
	}

	return allPrerequisites[course]
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)

	for _, pre := range prerequisites {
		adj[pre[0]] = append(adj[pre[0]], pre[1])
	}

	visit := make([]bool, numCourses)
	allPrerequisites := make([]map[int]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		dfs(adj, i, allPrerequisites, visit)
	}

	result := make([]bool, len(queries))

	for i, query := range queries {
		if allPrerequisites[query[0]][query[1]] {
			result[i] = true
		}
	}

	return result
}
