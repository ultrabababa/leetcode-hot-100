package permutations

func permute(nums []int) [][]int {
	var res [][]int
	var path []int

	// used 数组记录哪些元素已经被加入到了当前的 path 中
	used := make([]bool, len(nums))

	// 定义回溯函数
	var backtrack func()
	backtrack = func() {
		// 1. 结束条件：路径长度等于数组长度，说明找齐了一个排列
		if len(path) == len(nums) {
			// ⚠️ 极其重要的避坑点：切片是引用传递！
			// 必须深拷贝一份当前的 path，否则后续的回溯操作会把存进 res 里的数据改掉
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// 2. 遍历所有的选择
		for i := 0; i < len(nums); i++ {
			// 如果这个数字已经被用过了，跳过
			if used[i] {
				continue
			}

			// === 做选择 ===
			path = append(path, nums[i]) // 把数字加到路径里
			used[i] = true               // 标记为已使用

			// === 递归进入下一层决策树 ===
			backtrack()

			// === 撤销选择 (回溯) ===
			// 从下一层退回来时，要把刚刚加进去的数字拿出来，恢复现场
			// 这样在下一次 for 循环时，才能干干净净地去选别的数字
			path = path[:len(path)-1] // 弹出最后一个元素
			used[i] = false           // 标记为未使用
		}
	}

	// 从空路径开始触发回溯
	backtrack()

	return res
}
