package binary_tree_maximum_path_sum

import (
	. "leetcode/hot100/tree/common"
	"math"
)

func maxPathSum(root *TreeNode) int {
	// 初始化为极小值，因为节点值可能是负数
	maxSum := math.MinInt32

	// 定义 dfs 函数
	// 它的作用是：计算以 node 为起点的“单行道”最大向下收益
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		// Base Case
		if node == nil {
			return 0
		}

		// 1. 去左右子树打探“最大单边收益” (后序遍历)
		// 🚨 核心逻辑：如果收益是负数，直接抛弃（和 0 取最大值）
		leftGain := max(dfs(node.Left), 0)
		rightGain := max(dfs(node.Right), 0)

		// 2. 算总账：如果把当前节点当做“拱桥”的顶点，这桥的总和是多少？
		// 因为 leftGain 和 rightGain 已经保证非负了，所以放心加
		currentArchSum := node.Val + leftGain + rightGain

		// 3. 偷更全局最大值：有没有破历史记录？
		maxSum = max(maxSum, currentArchSum)

		// 4. 向上汇报：告诉父节点，我这条单行道能给你带来多少收益
		// 父节点只能挑我的一边往下走，所以我只能把较大的一边贡献上去
		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return maxSum
}
