// hot100/stack/739-daily-temperatures/solution.go

package main

func dailyTemperatures(temperatures []int) []int {
	// 单调栈，栈内存储的是下标, 维持栈底到栈顶的下标对应的元素单调递减
	var stack []int
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 || temperatures[i] <= temperatures[stack[len(stack)-1]] {
			// 如果栈为空或者当前温度小于等于栈顶温度，直接入栈，这样可以保证单调递减
			stack = append(stack, i)
		} else {
			// 否则栈不空且当前温度大于栈顶温度
			// 那就要循环比较栈顶元素，更新result，直到栈空或当前小于栈顶为止
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if temperatures[i] <= temperatures[top] {
					break
				}
				// 当前比栈顶的温度高，计算结果并出栈
				result[top] = i - top
				stack = stack[:len(stack)-1]
			}
			// 处理完了把i入栈
			stack = append(stack, i)
		}
		// 遍历完所有元素后，最后栈里剩下的元素不用处理，因为它们单调递减，result默认值为0
	}
	return result
}
