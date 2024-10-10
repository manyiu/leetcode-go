package sortitemsbygroupsrespectingdependencies

func topSortDfs(adjList [][]int, curr int, visit, path []bool, output *[]int) bool {
	if path[curr] {
		return false
	}

	if visit[curr] {
		return true
	}

	path[curr] = true
	visit[curr] = true

	for i := 0; i < len(adjList[curr]); i++ {
		if !topSortDfs(adjList, adjList[curr][i], visit, path, output) {
			return false
		}
	}

	*output = append(*output, curr)
	path[curr] = false

	return true
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	for i := 0; i < len(group); i++ {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	beforeGroups := make([][]int, m)

	for i := 0; i < len(beforeItems); i++ {
		for j := 0; j < len(beforeItems[i]); j++ {
			if group[i] != group[beforeItems[i][j]] {
				beforeGroups[group[i]] = append(beforeGroups[group[i]], group[beforeItems[i][j]])
			}
		}
	}

	itemPath := make([]bool, n)
	itemVisit := make([]bool, n)
	sortedItem := []int{}

	for i := 0; i < len(beforeItems); i++ {
		if !topSortDfs(beforeItems, i, itemVisit, itemPath, &sortedItem) {
			return []int{}
		}
	}

	groupPath := make([]bool, m)
	groupVisit := make([]bool, m)
	sortedGroup := []int{}

	for i := 0; i < len(beforeGroups); i++ {
		if !topSortDfs(beforeGroups, i, groupVisit, groupPath, &sortedGroup) {
			return []int{}
		}
	}

	itemsByGroup := make([][]int, m)

	for i := 0; i < n; i++ {
		itemsByGroup[group[sortedItem[i]]] = append(itemsByGroup[group[sortedItem[i]]], sortedItem[i])
	}

	result := make([]int, n)

	k := 0

	for i := 0; i < m; i++ {
		for j := 0; j < len(itemsByGroup[sortedGroup[i]]); j++ {
			result[k] = itemsByGroup[sortedGroup[i]][j]
			k++
		}
	}

	return result
}
