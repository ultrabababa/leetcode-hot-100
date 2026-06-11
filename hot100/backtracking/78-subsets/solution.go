package subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(startIndex int)
	backtrack = func(startIndex int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])

			backtrack(i + 1)

			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
