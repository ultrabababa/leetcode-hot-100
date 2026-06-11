package diameter_of_binary_tree

import . "leetcode/hot100/tree/common"

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int // 记录全局最大直径

	// 定义 dfs 函数
	// 这里的返回值是：以 node 为根的子树的【深度】（节点数）
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// Base Case
		if node == nil {
			return 0
		}

		// 1. 递归计算左右子树的深度
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)

		// 2. 核心逻辑：更新直径
		// 穿过当前节点的最长路径 = 左深度 + 右深度
		// 为什么不用 +1？因为题目定义直径是【边的数量】
		// 左边有 L 个节点深，右边有 R 个节点深，连起来就是 L+R 条边
		ans = max(ans, leftDepth+rightDepth)

		// 3. 返回当前节点的深度给父节点
		// 深度 = 左右子树中较深的那个 + 1 (自己)
		return max(leftDepth, rightDepth) + 1
	}

	dfs(root)
	return ans
}
