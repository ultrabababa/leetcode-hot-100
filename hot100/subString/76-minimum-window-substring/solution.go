package minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 1. 统计 t 中每个字符需要的数量
	// 使用 ASCII 数组代替 map，速度更快
	need := [128]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// left, right: 窗口左右指针
	// count: 还需要凑齐的字符总数 (比如 t="ABC", count=3; t="AA", count=2)
	// minLen: 记录遇到的最小窗口长度，初始化为无穷大
	// start: 记录最小窗口的起始位置
	left, right := 0, 0
	count := len(t)
	minLen := math.MaxInt32
	start := 0
	// 2. 滑动窗口开始
	for right < len(s) {
		ch := s[right]
		// --- 进窗口逻辑 ---
		// need[ch] > 0 代表这个字符是我们需要凑的
		// 如果是无关字符或者多余字符，need[ch] 会是 0 或负数
		if need[ch] > 0 {
			count--
		}
		// 无论需不需要，都把需求量减1
		// 如果变成负数，说明窗口里这个字符已经“超量”了，或者根本不需要
		need[ch]--
		right++

		// --- 出窗口逻辑 ---
		// 当 count == 0 时，说明窗口已经覆盖了 t 的所有字符
		// 我们尝试收缩左边界，看能不能让窗口变小一点
		// 更新最小窗口记录
		// 注意：此时 right 已经加了1，所以长度是 right - left
		for count == 0 {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			removedCh := s[left]
			// 如果加回去后 > 0，说明我们移除的是一个“必需品”
			// 移除它会导致窗口不再满足条件，所以 count++
			need[removedCh]++
			left++
			if need[removedCh] > 0 {
				count++
			}
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}
