package numberofislands

func mark(grid [][]byte, row int, column int) {
	if grid[row][column] == '1' {
		grid[row][column] = '2'

		if row > 0 {
			mark(grid, row-1, column)
		}

		if row < len(grid)-1 {
			mark(grid, row+1, column)
		}

		if column > 0 {
			mark(grid, row, column-1)
		}

		if column < len(grid[0])-1 {
			mark(grid, row, column+1)
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0

	for r, row := range grid {
		for c, e := range row {
			if e == '1' {
				count += 1
				mark(grid, r, c)
			}
		}
	}

	return count
}
