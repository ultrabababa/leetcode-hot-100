package trapping_rain_water

func trap(height []int) int {
	// 核心原理：按列求水 你可以想象自己站在第 i 根柱子上，你头顶能存多少水，取决于： min(左边最高的墙, 右边最高的墙) - 自己的高度
	// 前后缀最大值解法
	n := len(height)
	if n == 0 {
		return 0
	}

	preMax := make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		// 改进点 1: 直接使用内置 max 函数 (Go 1.21+)
		preMax[i] = max(preMax[i-1], height[i])
	}

	sufMax := make([]int, n)
	sufMax[n-1] = height[n-1]
	// 改进点 2: i >= 0 更符合阅读习惯
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		// 改进点 3: 逻辑合并，更加简洁
		// 这一列的水 = min(左边最高, 右边最高) - 自己的高度
		res += min(preMax[i], sufMax[i]) - height[i]
	}

	return res
}
