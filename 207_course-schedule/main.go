package courseschedule

func dfs(al map[int][]int, cur int, v map[int]bool) bool {
	if len(al[cur]) == 0 {
		return true
	}

	if v[cur] {
		return false
	}

	v[cur] = true

	for _, c := range al[cur] {
		if !dfs(al, c, v) {
			return false
		}
	}

	v[cur] = false
	al[cur] = []int{}

	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacencyList := map[int][]int{}
	visited := map[int]bool{}

	for _, prerequisite := range prerequisites {
		adjacencyList[prerequisite[0]] = append(adjacencyList[prerequisite[0]], prerequisite[1])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(adjacencyList, i, visited) {
			return false
		}
	}

	return true
}
