package first_missing_positive

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		// 核心逻辑：While 循环不断交换
		// 1. nums[i] > 0: 我们只关心正数
		// 2. nums[i] <= n: 数字不能越界 (比如数组长3，出现数字4我们是不管的，因为它占不到坑)
		// 3. nums[nums[i]-1] != nums[i]: 目标位置上的数字如果不等于当前数字，才交换
		//    (这一步既防止了死循环，也处理了重复元素，比如 [1, 1])
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 目标索引
			targetIndex := nums[i] - 1

			// 交换当前位置和目标位置的元素
			nums[i], nums[targetIndex] = nums[targetIndex], nums[i]
		}
	}

	// 再次遍历，寻找第一个“不在岗”的元素
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 如果所有位置都对上了 (例如 [1, 2, 3])
	// 那么缺失的就是下一个正数
	return n + 1
}
