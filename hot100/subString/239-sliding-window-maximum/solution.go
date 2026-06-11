package sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	// 双端队列，存储的是 nums 的 下标 (index)
	// 队列性质：nums[q[0]], nums[q[1]]... 是单调递减的
	q := make([]int, 0)

	for i, num := range nums {
		// 1. 【入队前清理】
		// 如果队列里“尾部”的元素比当前元素小，它们这就没用了，全部踢出
		// 因为当前元素 num 比它们大，而且比它们出现得晚，
		// 所以那些小的元素永远不可能成为窗口里的最大值了。
		for len(q) > 0 && num > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		// 2. 【新元素入队】
		q = append(q, i)
		// 3. 【检查过期】
		// 队头元素是当前最大值，但我们要看它还在不在窗口内
		// 窗口范围是 [i-k+1, i]，如果队头下标 < i-k+1，说明过期了
		if q[0] < i-k+1 {
			q = q[1:]
		}

		// 4. 【记录答案】
		// 当窗口长度达到 k 时，开始记录结果
		// (比如 k=3，当 i=2 时，窗口 [0,1,2] 形成，开始记录)
		if i >= k-1 {
			result = append(result, nums[q[0]])
		}
	}
	return result
}
