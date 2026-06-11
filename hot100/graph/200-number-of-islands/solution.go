package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0 // 记录岛屿数量

	// 定义 DFS 函数：职责是“沉没”以 (r, c) 为起点的整座岛屿
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// 1. 递归终止条件 (越界，或者是水)
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
			return
		}

		// 2. 核心动作：把陆地变成水 (沉岛)
		grid[r][c] = '0'

		// 3. 继续向上下左右四个方向蔓延
		dfs(r-1, c) // 上
		dfs(r+1, c) // 下
		dfs(r, c-1) // 左
		dfs(r, c+1) // 右
	}

	// 遍历整个网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 一旦发现一块陆地
			if grid[i][j] == '1' {
				count++   // 岛屿数量 +1
				dfs(i, j) // 立即启动 DFS，把这块陆地相连的整座岛炸沉
			}
		}
	}

	return count
}
