package convert_sorted_array_to_binary_search_tree

import . "leetcode/hot100/tree/common"

func sortedArrayToBST(nums []int) *TreeNode {
	// Base Case: 如果切片为空，说明没节点了
	if len(nums) == 0 {
		return nil
	}

	// 1. 找中点
	// 比如 len=5, mid=2 (index 2) -> 值是 0
	// 比如 len=2, mid=1 (index 1) -> 值是 3
	mid := len(nums) / 2

	// 2. 以中点为根构建节点
	root := &TreeNode{Val: nums[mid]}

	// 3. 递归构建左子树 (切片 nums[:mid] 不包含 mid)
	root.Left = sortedArrayToBST(nums[:mid])

	// 4. 递归构建右子树 (切片 nums[mid+1:] 从 mid 后面开始)
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
