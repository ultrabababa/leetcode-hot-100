package threesum

import "sort"

func threeSum(nums []int) [][]int {
	// 排序 (Sorting)：先用 sort.Ints(nums) 把数组排好序。
	// 这一步至关重要，它把无序的问题变成了有序的问题，让去重变得简单。
	sort.Ints(nums)
	// 固定一个，找另外两个：
	//我们需要 a + b + c = 0。
	//我们遍历数组，固定第一个数 nums[i]。
	//问题就变成了：在 i 后面的数组里，找到两个数 L 和 R，使得 nums[L] + nums[R] == -nums[i]。
	//这就转化成了我们刚做过的 两数之和 或者类似 盛水容器 的双指针问题。
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		num := nums[i]
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		// 双指针移动：
		//L 在 i+1，R 在最右边。
		//若和 > 0，说明太大了，R 往左移。
		//若和 < 0，说明太小了，L 往右移。
		//若和 == 0，记录答案，并且 L, R 同时收缩。
		for left < right {
			// 最难点：如何去重？（Bug 高发区）
			//题目说不准有重复的三元组。因为数组排过序了，重复的元素一定靠在一起。
			//针对 i 去重： 如果 nums[i] == nums[i-1]，说明这个数刚才已经当过“带头大哥”了，再用它当头，找出来的结果肯定一样。直接跳过 (continue)。
			//针对 L 和 R 去重： 当我们找到一组解 [-1, -1, 2] 后，L 指针如果向右移一步还是 -1，那找出来的解肯定还是重复的。 所以，找到答案后，要循环跳过所有相同的值。
			sum := num + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
