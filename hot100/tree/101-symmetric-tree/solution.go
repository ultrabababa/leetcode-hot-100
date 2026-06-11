package symmetric_tree

import . "leetcode/hot100/tree/common"

// 递归法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 启动双指针递归：比较左子树和右子树
	return check(root.Left, root.Right)
}

func check(p, q *TreeNode) bool {
	// 1. 两个都为空：是对称的
	if p == nil && q == nil {
		return true
	}
	// 2. 只有一个为空：不对称
	if p == nil || q == nil {
		return false
	}
	// 3. 值不相等：不对称
	if p.Val != q.Val {
		return false
	}

	// 4. 递归检查：
	// p 的左边 vs q 的右边 (外侧)
	// p 的右边 vs q 的左边 (内侧)
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 迭代法
func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 初始化队列，放入这一对需要比较的节点
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		// 一次拿出两个
		u, v := queue[0], queue[1]
		queue = queue[2:]

		// 1. 都为空，这一枝是对称的，继续看别的
		if u == nil && v == nil {
			continue
		}
		// 2. 一个为空，或者值不等，直接死刑
		if u == nil || v == nil || u.Val != v.Val {
			return false
		}

		// 3. 成对加入子节点 (顺序很重要！)
		// 比较 u.Left 和 v.Right
		queue = append(queue, u.Left, v.Right)
		// 比较 u.Right 和 v.Left
		queue = append(queue, u.Right, v.Left)
	}

	return true
}
