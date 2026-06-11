package letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}

	phoneMap := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	var res []string
	var path []byte

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		digit := digits[index] - '0'
		letters := phoneMap[digit]

		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])

			backtrack(index + 1)

			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
