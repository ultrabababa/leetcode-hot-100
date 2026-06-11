package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 1. 先计算【左侧乘积】
	// answer[i] 表示 i 左边所有元素的乘积
	// 对于索引 0，左边没有元素，所以初始化为 1
	answer[0] = 1
	for i := 1; i < n; i++ {
		// answer[i] = (i-1 左边的乘积) * nums[i-1]
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 2. 再计算【右侧乘积】并累乘
	// R 为右侧所有元素的乘积，初始为 1
	R := 1
	for i := n - 1; i >= 0; i-- {
		// 此时 answer[i] 已经是左侧乘积了
		// 我们再乘以右侧乘积 R，就是最终结果
		answer[i] = answer[i] * R

		// 更新 R，让它包含当前的 nums[i]，为下一次循环（更左边的位置）做准备
		R *= nums[i]
	}

	return answer
}
