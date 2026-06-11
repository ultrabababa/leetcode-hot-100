package word_search

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	// 定义 DFS 函数
	// r, c: 当前所在的网格坐标
	// index: 当前需要匹配 word 中的第几个字母
	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		// Base Case 1：单词全部匹配完了，探险成功！
		if index == len(word) {
			return true
		}

		// Base Case 2：越界，或者当前格子的字母不匹配，或者遇到了走过的格子 ('#')
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[index] {
			return false
		}

		// === 做选择 ===
		// 暂存当前字符，并把当前格子标记为已访问 (沉岛)
		temp := board[r][c]
		board[r][c] = '#'

		// === 递归进入下一层 ===
		// 向上下左右四个方向探索下一个字母
		// 只要有一个方向能走通 (返回 true)，我们就成功了
		found := dfs(r-1, c, index+1) || // 上
			dfs(r+1, c, index+1) || // 下
			dfs(r, c-1, index+1) || // 左
			dfs(r, c+1, index+1) // 右

		// === 撤销选择 (回溯) ===
		// 探索完毕，无论成败，都要把当前格子恢复原样，留给平行宇宙的兄弟去试
		board[r][c] = temp

		return found
	}

	// 遍历整个网格，寻找可能的起点
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 只有首字母匹配，才值得启动 DFS 探险队
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true // 只要找到一条路径就立刻下班
				}
			}
		}
	}

	return false
}
