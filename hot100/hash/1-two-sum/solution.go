package twosum

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if preIndex, ok := indexMap[complement]; ok {
			// found
			return []int{preIndex, i}
		}
		indexMap[num] = i
	}
	return nil
}
