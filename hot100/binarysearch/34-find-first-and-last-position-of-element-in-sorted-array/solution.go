package main

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)

	if first == -1 {
		// 没找到
		return []int{-1, -1}
	}
	last := findLast(nums, target)

	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	// 找到左边界
	left, right := 0, len(nums)-1
	first := -1

	// 二分查找不停找，最后first就是最左边的target的位置
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			first = mid
			right = mid - 1 // 向左收缩继续找，而不是return
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return first
}

func findLast(nums []int, target int) int {
	// 找到左边界
	left, right := 0, len(nums)-1
	last := -1

	// 二分查找不停找，最后first就是最左边的target的位置
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			last = mid
			left = mid + 1 // 向右收缩继续找，而不是return
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return last
}
