package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SliceToTree ==========================================
// 核心工具：SliceToTree (数组 -> 二叉树)
// ==========================================
// 输入：[]interface{}{1, nil, 2, 3}
// 输出：树的根节点
// 注意：必须使用 interface{}，因为 Go 的 int 不能为 nil
func SliceToTree(data []interface{}) *TreeNode {
	if len(data) == 0 || data[0] == nil {
		return nil
	}

	// 1. 创建根节点
	root := &TreeNode{Val: data[0].(int)}

	// 2. 使用队列辅助构建 (这是标准的层序遍历逻辑)
	queue := []*TreeNode{root}
	i := 1 // 数组索引，从 1 开始（0是根节点）

	for i < len(data) {
		// 取出当前父节点
		parent := queue[0]
		queue = queue[1:]

		// 处理左孩子
		if i < len(data) {
			val := data[i]
			if val != nil {
				leftNode := &TreeNode{Val: val.(int)}
				parent.Left = leftNode
				queue = append(queue, leftNode)
			}
			i++
		}

		// 处理右孩子
		if i < len(data) {
			val := data[i]
			if val != nil {
				rightNode := &TreeNode{Val: val.(int)}
				parent.Right = rightNode
				queue = append(queue, rightNode)
			}
			i++
		}
	}

	return root
}

// TreeToSlice ==========================================
// 辅助工具：TreeToSlice (二叉树 -> 数组)
// ==========================================
// 用于在测试中断言结果，将树转换回层序遍历的数组
// 例如：[1, 2, 3, nil, nil, 4, 5]
func TreeToSlice(root *TreeNode) []int {
	// 很多题目（如中序遍历）返回的是特定顺序的数组，
	// 所以这个函数主要用于调试打印，或者测试层序遍历题目。
	// 如果你想验证中序遍历结果，直接对比题目要求的 []int 即可。

	// 这里为了简单，只演示层序遍历的转换逻辑
	if root == nil {
		return []int{}
	}
	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			// 这里有个小问题：LeetCode 的输出通常会省略末尾的 null
			// 但为了简单调试，我们可以用 -1 或特定值代表 null，或者不存
		}
	}
	return res
}
