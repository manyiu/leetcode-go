package numberofislands

func dfs(grid [][]byte, r, c int) {
	if grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0'

	dir := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	for _, step := range dir {
		if r+step[0] >= 0 && r+step[0] < len(grid) && c+step[1] >= 0 && c+step[1] < len(grid[r]) {
			dfs(grid, r+step[0], c+step[1])
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0

	for i, row := range grid {
		for j, column := range row {
			if column == '1' {
				count++

				dfs(grid, i, j)
			}
		}
	}

	return count
}
