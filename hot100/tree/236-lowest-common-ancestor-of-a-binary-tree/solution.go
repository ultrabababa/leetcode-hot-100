package lowest_common_ancestor_of_a_binary_tree

import . "leetcode/hot100/tree/common"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 1. Base Case (递归终止条件)
	// 如果越过叶子节点，返回 nil
	// 如果当前节点就是 p 或者 q，说明找到了目标，立刻向上返回该节点
	if root == nil || root == p || root == q {
		return root
	}

	// 2. 去左子树和右子树里“打探情报” (后序遍历)
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 3. 汇总情报，做出判断
	// 情况 A: 左右子树分别传回来了 p 和 q (都不为空)
	// 说明 p 和 q 分居在当前节点的两侧，当前节点就是最近公共祖先！
	if left != nil && right != nil {
		return root
	}

	// 情况 B: 左子树没找到，说明 p 和 q 都在右子树
	// 此时 right 里面存的要么是 p，要么是 q，要么已经是它俩的 LCA，直接传给上层
	if left == nil {
		return right
	}

	// 情况 C: 右子树没找到，说明 p 和 q 都在左子树
	if right == nil {
		return left
	}

	// 理论上不会走到这里，因为题目保证 p 和 q 都在树里
	return nil
}
