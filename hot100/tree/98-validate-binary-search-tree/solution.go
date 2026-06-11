package validate_binary_search_tree

import (
	. "leetcode/hot100/tree/common"
)

// 利用 BST 的一个绝对性质：
//“二叉搜索树的中序遍历结果，一定是一个严格递增的序列。”
//
//我们不需要把整个数组存下来再检查。我们只需要在遍历过程中，记录前一个访问的节点值 (pre)。
//只要发现 当前节点值 <= 前一个节点值，就说明顺序乱了，直接返回 false。

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) {
			return false
		}
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		if !dfs(node.Right) {
			return false
		}
		return true
	}
	return dfs(root)
}
