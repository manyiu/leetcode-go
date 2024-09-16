package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	grid := make([][]int, len(obstacleGrid))

	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		row := make([]int, len(obstacleGrid[i]))

		for j := len(obstacleGrid[i]) - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				row[j] = 0
			} else if i == len(obstacleGrid)-1 && j == len(obstacleGrid[i])-1 {
				row[j] = 1
			} else if i == len(obstacleGrid)-1 {
				row[j] = row[j+1]
			} else if j == len(obstacleGrid[i])-1 {
				row[j] = grid[i+1][j]
			}
		}

		grid[i] = row
	}

	for j := len(grid) - 2; j >= 0; j-- {
		for k := len(grid[j]) - 2; k >= 0; k-- {
			if obstacleGrid[j][k] == 1 {
				grid[j][k] = 0
			} else {
				grid[j][k] = grid[j+1][k] + grid[j][k+1]
			}
		}
	}

	return grid[0][0]
}
