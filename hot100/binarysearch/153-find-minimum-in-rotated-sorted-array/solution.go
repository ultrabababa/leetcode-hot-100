// hot100/binarysearch/153-find-minimum-in-rotated-sorted-array/solution.go

package main

import "math"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	minNum := math.MaxInt

	for left <= right {
		mid := left + (right-left)/2

		if nums[left] <= nums[mid] {
			// 左半严格增，取其中的最小值后转到右半继续找更小值
			if nums[left] < minNum {
				minNum = nums[left]
			}
			left = mid + 1
		} else {
			// 右半严格增，取其中的最小值nums[mid]后到左半找更小值
			if nums[mid] < minNum {
				minNum = nums[mid]
			}
			right = mid - 1
		}
	}

	return minNum
}
