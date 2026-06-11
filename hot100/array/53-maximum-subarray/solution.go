package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// pre 代表以“前一个元素”结尾的最大子数组和
	// maxAns 代表全局已知的最大和
	// 初始化为第一个元素，防止全负数时出错
	preSum := nums[0]
	maxAns := nums[0]

	for i := 1; i < len(nums); i++ {
		// 决策核心：
		// 如果前面的累积和 (pre) 大于 0，我们就带上它 (pre + nums[i])
		// 如果前面的累积和 <= 0，它是累赘，我们就扔掉它，从自己重新开始 (nums[i])
		if preSum > 0 {
			preSum += nums[i]
		} else {
			preSum = nums[i]
		}
		// 每次更新完当前状态，都要挑战一下全局最大值
		if preSum > maxAns {
			maxAns = preSum
		}
	}
	return maxAns
}
