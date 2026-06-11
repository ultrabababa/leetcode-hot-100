package palindrome_partitioning

func partition(s string) [][]string {
	var res [][]string
	var path []string

	isPalindrome := func(str string, left, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	var backtrack func(startIndex int)
	backtrack = func(startIndex int) {
		if startIndex == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if isPalindrome(s, startIndex, i) {
				path = append(path, s[startIndex:i+1])

				backtrack(i + 1)

				path = path[:len(path)-1]
			} else {
				continue
			}
		}
	}
	backtrack(0)
	return res
}
