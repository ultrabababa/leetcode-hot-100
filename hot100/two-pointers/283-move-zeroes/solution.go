package move_zeroes

func moveZeroes(nums []int) {
	// 题目本质就是把非零元素都挪到数组前面去
	slow := 0 // 永远指向下一个准备接受非零值的位置
	for fast := 0; fast < len(nums); fast++ {
		// fast用来寻找非零元素
		if nums[fast] != 0 {
			// 如果找到了非零元素就和fast的值交换
			// 判断一下如果fast和slow本身已经是同一个位置了，就不用交换减少操作
			if fast != slow {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			// slow之前都是非零元素了
			slow++
		}
	}
}
