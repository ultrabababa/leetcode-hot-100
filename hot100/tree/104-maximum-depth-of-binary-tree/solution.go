package maximum_depth_of_binary_tree

import . "leetcode/hot100/tree/common"

func maxDepth(root *TreeNode) int {
	// Base Case: 空节点深度为 0
	if root == nil {
		return 0
	}

	// 1. 计算左子树深度
	leftDepth := maxDepth(root.Left)

	// 2. 计算右子树深度
	rightDepth := maxDepth(root.Right)

	// 3. 取最大值 + 1
	return max(leftDepth, rightDepth) + 1
}
