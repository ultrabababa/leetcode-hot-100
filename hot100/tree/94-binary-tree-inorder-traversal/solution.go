package binary_tree_inorder_traversal

import . "leetcode/hot100/tree/common"

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		// 1. 递归终止条件
		if node == nil {
			return
		}

		// 2. 按照 左 -> 根 -> 右 的顺序
		dfs(node.Left)              // 左
		res = append(res, node.Val) // 根
		dfs(node.Right)             // 右
	}

	dfs(root)
	return res
}

// 迭代法
func inorderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0)
	var stack []*TreeNode
	cur := root

	// 循环条件：
	// 1. cur != nil: 说明还有新节点没探索（通常是刚转向右子树时）
	// 2. len(stack) > 0: 说明栈里还有“父皇”在等着被“解冻”
	for cur != nil || len(stack) > 0 {

		// 【阶段一：疯狂向左】
		if cur != nil {
			stack = append(stack, cur) // 记录路过的节点（暂不处理值）
			cur = cur.Left             // 继续向左冲
		} else {
			// 【阶段二：撞墙回头】
			// cur 已经是 nil 了，说明左边没路了

			// 1. 回溯：从栈顶拿出最近的一个父节点
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop

			// 2. 办事：中序遍历在这里记录值
			res = append(res, node.Val)

			// 3. 转向：左边和根都搞定了，该去右边看看了
			// 此时 cur 指向了右孩子。
			// 如果右孩子存在，下一轮循环会进入【阶段一】继续向左。
			// 如果右孩子不存在(nil)，下一轮循环会再次进入【阶段二】，继续往上回溯。
			cur = node.Right
		}
	}
	return res
}
