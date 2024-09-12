package shortestpathinbinarymatrix

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	queue := [][]int{{0, 0}}
	length := 0

	dir := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for len(queue) > 0 {
		length++
		for _, cell := range queue {
			queue = queue[1:]

			if grid[cell[0]][cell[1]] == 1 {
				continue
			}

			grid[cell[0]][cell[1]] = 1

			if cell[0] == len(grid)-1 && cell[1] == len(grid[cell[0]])-1 {
				return length
			}

			for _, step := range dir {
				if cell[0]+step[0] >= 0 &&
					cell[0]+step[0] < len(grid) &&
					cell[1]+step[1] >= 0 &&
					cell[1]+step[1] < len(grid[cell[0]]) &&
					grid[cell[0]+step[0]][cell[1]+step[1]] == 0 {

					queue = append(queue, []int{cell[0] + step[0], cell[1] + step[1]})
				}
			}
		}
	}

	return -1
}
