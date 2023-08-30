package matrix

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	queue := make([][]int, 0)

	for i, row := range mat {
		for j, val := range row {
			if val == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			x := point[0] + dir[0]
			y := point[1] + dir[1]

			if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == -1 {
				mat[x][y] = mat[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}

	return mat
}
