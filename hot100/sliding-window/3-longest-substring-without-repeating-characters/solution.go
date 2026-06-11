package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	// 核心思路提示
	//我们需要维护一个窗口 [left, right]，保证这个窗口里的字符 全部不重复。
	//扩大窗口： right 主动向右走，把新字符加进来。
	//收缩窗口（关键）：
	//如果 right 指向的新字符，在窗口里已经出现过了（比如叫 'a'），说明窗口不再满足“无重复”了。
	//这时候，我们需要移动 left，把窗口左边的元素踢出去，直到 把那个重复的 'a' 踢出去为止。
	//优化（Map）： 我们可以用一个 Map 记录每个字符上一次出现的下标。当遇到重复字符时，left 可以直接“跳”到那个字符上一次出现位置的下一位(lastPos+1)，而不用一步步挪。
	// 1. 哈希表：记录字符上一次出现的位置
	// key: 字符 (byte), value: 下标 (int)
	lastOccurred := make(map[byte]int)
	start := 0
	maxRes := 0

	for end, ch := range []byte(s) {
		if lastPos, ok := lastOccurred[ch]; ok {
			start = max(start, lastPos+1)
		}
		lastOccurred[ch] = end
		maxRes = max(maxRes, end-start+1)
	}
	return maxRes
}
