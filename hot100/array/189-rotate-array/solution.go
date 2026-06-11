package rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// 1. 预处理 k
	// 如果 k > n，我们只需要转 k % n 次
	// 比如 n=5, k=7，转7次其实就是转2次
	k = k % n

	// 2. 三次翻转法

	// 第一步：翻转全部
	reverse(nums, 0, n-1)

	// 第二步：翻转前 k 个 (索引 0 到 k-1)
	reverse(nums, 0, k-1)

	// 第三步：翻转剩下的 (索引 k 到 n-1)
	reverse(nums, k, n-1)
}

// 辅助函数：翻转数组 nums 中从 start 到 end 的部分
// 双指针法，O(N) 时间，O(1) 空间
func reverse(nums []int, start, end int) {
	for start < end {
		// 交换
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
