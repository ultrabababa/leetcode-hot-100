package search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	// 从右上角开始搜索
	row := 0
	col := n - 1

	// 只要没有越界，就一直找
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true // 找到了
		} else if matrix[row][col] > target {
			// 当前值太大，因为列是从上到下递增的，所以这一列下面更大
			// 只能往左找更小的
			col--
		} else {
			// 当前值太小，因为行是从左到右递增的，这一行左边更小
			// 只能往下找更大的
			row++
		}
	}

	// 越界了还没找到
	return false
}
