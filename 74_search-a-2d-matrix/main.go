package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	t, b := 0, len(matrix)-1

	for t < b {
		m := (t + b) / 2

		if target >= matrix[m][0] && target <= matrix[m][len(matrix[m])-1] {
			t, b = m, m
		}

		if target < matrix[m][0] {
			b = m - 1
		}

		if target > matrix[m][len(matrix[m])-1] {
			t = m + 1
		}
	}

	l, r := 0, len(matrix[t])-1

	for l <= r {
		m := (l + r) / 2

		if matrix[t][m] == target {
			return true
		}

		if matrix[t][m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
