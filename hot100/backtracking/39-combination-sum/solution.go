package combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	// ⭐️ 优化点 1：将数组从小到大排序，为后面的“剪枝”做准备
	sort.Ints(candidates)

	// 定义回溯函数
	// remain: 离凑齐目标和还差多少
	// startIndex: 当前循环从哪里开始挑选
	var backtrack func(remain, startIndex int)
	backtrack = func(remain, startIndex int) {
		// Base Case 1：刚好凑齐
		if remain == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// 遍历候选数字，从 startIndex 开始，保证不回头产生重复组合
		for i := startIndex; i < len(candidates); i++ {
			// ⭐️ 优化点 2 (剪枝)：
			// 如果连当前的候选数都比 remain 大，因为数组已排序，
			// 后面的数字只会更大，绝对装不下，直接结束当前层的循环！
			if remain-candidates[i] < 0 {
				break
			}

			// === 做选择 ===
			path = append(path, candidates[i])

			// === 递归进入下一层 ===
			// ⚠️ 核心所在：这里传的是 i，而不是 i+1！
			// 因为可以无限次重复选取当前的数字
			backtrack(remain-candidates[i], i)

			// === 撤销选择 (回溯) ===
			path = path[:len(path)-1]
		}
	}

	// 初始状态：距离目标还差 target，从索引 0 开始挑选
	backtrack(target, 0)

	return res
}
