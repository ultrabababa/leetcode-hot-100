package invert_binary_tree

import . "leetcode/hot100/tree/common"

func invertTree(root *TreeNode) *TreeNode {
	// 1. 递归终止条件
	if root == nil {
		return nil
	}

	// 2. 核心逻辑：交换左右子树
	// 这里利用了 Go 语言的平行赋值，同时完成了递归和交换
	// root.Left 变成了 翻转后的右子树
	// root.Right 变成了 翻转后的左子树
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	// 3. 返回当前根节点
	return root
}
