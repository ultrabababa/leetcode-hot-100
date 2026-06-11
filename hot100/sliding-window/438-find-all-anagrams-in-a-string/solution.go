package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	// 题目只包含小写字母。
	//我们可以用一个长度为 26 的数组 [26]int 来代替 Map。
	//index 0 代表 'a'，index 1 代表 'b'...
	//比较两个长度为 26 的整数数组是否相等，在 Go 语言里非常快（CPU 可以在很少的指令周期内完成）。
	// 算法流程：
	//特判： 如果 len(s) < len(p)，直接返回空。
	//初始化：
	//统计 p 的字符计数 pCount。
	//统计 s 中前 len(p) 个字符的计数 windowCount。
	//初次比较： 如果 pCount == windowCount，记录索引 0。
	//开始滑动： 从 i = 0 开始滑到 len(s) - len(p)。
	//进一个： 窗口右边吃进来一个新字符，计数 +1。
	//出一个： 窗口左边吐出去一个旧字符，计数 -1。
	//比较： 看看现在的 windowCount 和 pCount 是否相等。
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return []int{}
	}
	result := make([]int, 0)
	var pCount [26]int
	var windowCount [26]int

	for i, ch := range []byte(p) {
		pCount[ch-'a']++
		windowCount[s[i]-'a']++
	}

	if pCount == windowCount {
		result = append(result, 0)
	}

	for i := 0; i < sLen-pLen; i++ {
		leftCharIndex := s[i] - 'a'
		rightCharIndex := s[i+pLen] - 'a'
		windowCount[leftCharIndex]--
		windowCount[rightCharIndex]++
		if windowCount == pCount {
			result = append(result, i+1)
		}
	}
	return result
}
