package rotting_oranges

// 定义坐标结构体，方便在队列中存取
type point struct {
	r int
	c int
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	queue := []point{}
	freshCount := 0

	// 1. 初始化：统计新鲜橘子，将所有腐烂橘子加入队列
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// 如果一开始就没有新鲜橘子，需要 0 分钟
	if freshCount == 0 {
		return 0
	}

	minutes := 0
	// 定义四个方向：上、下、左、右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 2. 开始 BFS（只要队列有橘子，且还有新鲜橘子没被感染）
	// 注意：条件带上 freshCount > 0 可以防止最后一批橘子腐烂后，多算 1 分钟
	for len(queue) > 0 && freshCount > 0 {
		minutes++

		// 记录当前这一层（这一分钟）有多少个腐烂橘子要向外蔓延
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			// 出队
			curr := queue[0]
			queue = queue[1:]

			// 向四个方向感染
			for _, d := range dirs {
				nr, nc := curr.r+d[0], curr.c+d[1]

				// 如果没有越界，并且遇到了新鲜橘子
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					// 感染它！
					grid[nr][nc] = 2
					freshCount--
					// 将新腐烂的橘子加入队列，它将在下一分钟发威
					queue = append(queue, point{nr, nc})
				}
			}
		}
	}

	// 3. 结算：检查是否还有幸存的新鲜橘子
	if freshCount > 0 {
		return -1
	}
	return minutes
}
