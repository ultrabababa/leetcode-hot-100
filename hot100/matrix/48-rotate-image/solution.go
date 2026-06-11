package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)

	// 1. 转置 (Transpose)
	// 沿着对角线交换元素
	// 注意 j 从 i 开始遍历，只遍历对角线右上方区域，避免重复交换
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. 左右翻转每一行 (Reverse each row)
	for i := 0; i < n; i++ {
		// 双指针法翻转一行
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
