package uniquepaths

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)

	for i := 0; i < m; i++ {
		row := make([]int, n)

		for j := 0; j < n; j++ {
			if i == m-1 || j == n-1 {
				row[j] = 1
			}
		}

		grid[i] = row
	}

	for j := m - 2; j >= 0; j-- {
		for k := n - 2; k >= 0; k-- {
			grid[j][k] = grid[j+1][k] + grid[j][k+1]
		}
	}

	return grid[0][0]
}
