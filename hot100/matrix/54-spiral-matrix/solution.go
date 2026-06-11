package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// 1. 初始化四个边界
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	// 结果集，容量预估为总元素个数
	res := make([]int, 0, (bottom+1)*(right+1))

	for {
		// --- 1. 向右移动 (Left -> Right) ---
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// 上边界收缩
		top++
		// 检查：如果上边界超过下边界，说明走完了
		if top > bottom {
			break
		}

		// --- 2. 向下移动 (Top -> Bottom) ---
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 右边界收缩
		right--
		// 检查
		if left > right {
			break
		}

		// --- 3. 向左移动 (Right -> Left) ---
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		// 下边界收缩
		bottom--
		// 检查
		if top > bottom {
			break
		}

		// --- 4. 向上移动 (Bottom -> Top) ---
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		// 左边界收缩
		left++
		// 检查
		if left > right {
			break
		}
	}

	return res
}
