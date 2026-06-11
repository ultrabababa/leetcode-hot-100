package subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	// 2. 解题思路：前缀和 + 哈希表核心公式推导：定义 pre[i] 为 nums[0]...nums[i] 的和（前缀和）。
	//那么，子数组 nums[j...i] 的和可以表示为：$$Sum(j, i) = pre[i] - pre[j-1]$$题目要求子数组和为 $k$，
	//即：$$pre[i] - pre[j-1] = k$$
	//移项（最关键的一步）：$$pre[j-1] = pre[i] - k$$
	//算法逻辑：我们在遍历数组计算当前前缀和 pre[i] 时，只需要回头问一下历史记录：
	//“嘿，以前有没有出现过前缀和等于 pre[i] - k 的情况？
	//如果有，出现了几次？”如果以前出现过 $N$ 次 pre[i] - k，
	//那就意味着有 $N$ 个子数组是以当前位置结尾且和为 $k$ 的。
	res := 0
	// key: 前缀和, value: 该前缀和出现的次数
	preSumMap := make(map[int]int)
	preSumMap[0] = 1
	currPreSum := 0
	for _, num := range nums {
		currPreSum += num
		target := currPreSum - k

		if times, ok := preSumMap[target]; ok {
			res += times
		}
		preSumMap[currPreSum]++
	}
	return res
}
