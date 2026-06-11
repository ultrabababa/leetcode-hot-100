package n_queens

func solveNQueens(n int) [][]string {
	var res [][]string

	// 初始化棋盘，全填充为 '.'
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	// 记录列和两条对角线的占用情况
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n) // 主对角线 \ : row - col 是常数，加 n 避免负数
	diag2 := make([]bool, 2*n) // 副对角线 / : row + col 是常数

	// 定义回溯函数，每次处理第 row 行
	var backtrack func(row int)
	backtrack = func(row int) {
		// Base Case：如果所有的行都处理完了，说明成功放置了 n 个皇后
		if row == n {
			// 把当前 board 的状态转成 []string 并存入结果集
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = string(board[i])
			}
			res = append(res, temp)
			return
		}

		// 尝试在当前行 (row) 的每一列 (col) 放置皇后
		for col := 0; col < n; col++ {
			// 计算当前格子对应的两条对角线索引
			d1 := row - col + n
			d2 := row + col

			// 剪枝：如果列或对角线已经被占用了，说明这个位置不安全，跳过
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			// === 做选择 ===
			board[row][col] = 'Q'
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			// === 递归进入下一层 (下一行) ===
			backtrack(row + 1)

			// === 撤销选择 (回溯) ===
			// 把皇后拿走，并且释放列和对角线的控制权
			board[row][col] = '.'
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	// 从第 0 行开始放皇后
	backtrack(0)

	return res
}
