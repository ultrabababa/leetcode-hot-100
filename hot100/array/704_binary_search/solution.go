package binarysearch

func Search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	// 因为我们初始定义 high = len - 1（闭区间）。如果数组只有一个元素 [1]，low=0, high=0。
	// 如果用 <，循环根本进不去，直接返回 -1，就错了。用 <= 才能检查这最后一个元素。
	for low <= high {
		mid := (low + high) / 2
		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			low = mid + 1
		case target < nums[mid]:
			high = mid - 1
		}
	}
	return -1
}
