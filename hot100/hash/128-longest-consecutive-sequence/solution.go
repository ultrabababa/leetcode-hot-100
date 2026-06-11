package longestconsecutive

// 核心思路提示：去重：既然 [1, 0, 1, 2] 算作 0, 1, 2 (长度3)，说明我们第一步应该把所有数字扔进一个 map[int]bool
// 里，既能去重，又能让我们用 $O(1)$ 的时间查某个数是否存在。找起点：这是最关键的一步。遍历数组中的每个数 x。如果 x - 1
// 存在于 map 中，说明 x 肯定不是一个序列的开头（因为它前面还有人），我们可以直接跳过它，不需要计算。只有当 x - 1
// 不在 map 中时，说明 x 是一个序列的新起点。向后数：一旦找到了起点 x，我们就疯狂地往后查 x+1, x+2, x+3
// 是否在 map 中，直到断掉为止。记录这个长度。

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}
	maxLen := 0
	for num, _ := range numsMap {
		if numsMap[num-1] {
			continue
		} else {
			currLen := 0
			for i := 0; numsMap[num+i]; i++ {
				currLen++
			}
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}
