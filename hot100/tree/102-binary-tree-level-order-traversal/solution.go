package binary_tree_level_order_traversal

import . "leetcode/hot100/tree/common"

func levelOrder(root *TreeNode) [][]int {
	// 1. 处理空树边界情况
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	// 2. 初始化队列
	queue := []*TreeNode{root}

	// 3. 开始 BFS 循环
	for len(queue) > 0 {
		// 【关键步骤】：记录当前层的节点数量 (快照)
		// 这一刻，队列里全是第 k 层的节点
		currentLevelSize := len(queue)
		var currentLevelVals []int

		// 4. 遍历当前层的所有节点
		for i := 0; i < currentLevelSize; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]

			// 收集值
			currentLevelVals = append(currentLevelVals, node.Val)

			// 将下一层节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 5. 将这一层的结果加入总结果集
		res = append(res, currentLevelVals)
	}

	return res
}
