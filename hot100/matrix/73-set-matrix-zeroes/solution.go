package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowZero := false
	colZero := false

	// 1. 检查第一列是否有 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}

	// 2. 检查第一行是否有 0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	// 3. 使用第一行和第一列作为标记位
	// 从 [1][1] 开始遍历，如果 matrix[i][j] 是 0，
	// 就把对应的行首 matrix[i][0] 和列首 matrix[0][j] 设为 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 4. 根据标记，将内部元素置为 0
	// 同样从 [1][1] 开始，看自己的行首或列首是不是 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 5. 最后处理第一行和第一列
	// 必须最后处理，否则会影响第 4 步的判断
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
