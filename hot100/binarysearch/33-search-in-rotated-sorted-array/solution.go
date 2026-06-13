// hot100/binarysearch/33-search-in-rotated-sorted-array/solution.go

package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断左半边是否是严格递增
		if nums[left] <= nums[mid] {
			// 判断target是否在左半部份的严格递增区间里
			if nums[left] <= target && nums[mid] > target {
				// 如果是那就缩小范围到这里面
				right = mid - 1
			} else {
				// 如果不是就排除这个区间
				left = mid + 1
			}
		} else {
			// 否则右半部份一定严格递增
			// 判断target是否在右半部份到严格递增区间里
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
