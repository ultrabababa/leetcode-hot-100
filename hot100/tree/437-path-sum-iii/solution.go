package path_sum_iii

import . "leetcode/hot100/tree/common"

func pathSum(root *TreeNode, targetSum int) int {
	// prefixMap 记录【前缀和】出现的【次数】
	// Key: 前缀和, Value: 出现的次数
	prefixMap := make(map[int]int)

	// 初始化边界条件：前缀和为 0 出现了一次
	// 意思是：如果某个节点的 currentSum 刚好等于 targetSum，
	// 那么 currentSum - targetSum = 0，我们需要能够查到这个 0。
	prefixMap[0] = 1

	var dfs func(node *TreeNode, currentSum int) int
	dfs = func(node *TreeNode, currentSum int) int {
		if node == nil {
			return 0
		}

		// 1. 累加得到到达当前节点的前缀和
		currentSum += node.Val

		// 2. 查表：看看之前有没有出现过 currentSum - targetSum 的前缀
		// 如果有，说明从那个节点到当前节点刚好是一条满足条件的路径
		count := prefixMap[currentSum-targetSum]

		// 3. 将当前的前缀和加入 Map，为后续的子节点做准备
		prefixMap[currentSum]++

		// 4. 继续往下钻，把左右子树里找到的路径数加起来
		count += dfs(node.Left, currentSum)
		count += dfs(node.Right, currentSum)

		// 5. 【核心】：状态恢复（回溯）
		// 离开当前节点，返回上一层时，必须把当前节点的前缀和次数减去。
		// 因为它不属于旁系分支（如它的兄弟节点）的路径前缀。
		prefixMap[currentSum]--

		return count
	}

	return dfs(root, 0)
}
