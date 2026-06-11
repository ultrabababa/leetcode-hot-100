package generate_parentheses

func generateParenthesis(n int) []string {
	var res []string

	// 定义回溯函数
	// left: 已经使用的左括号数量
	// right: 已经使用的右括号数量
	// path: 当前生成的括号字符串
	var backtrack func(left int, right int, path string)
	backtrack = func(left int, right int, path string) {
		// Base Case：如果长度够了 (n对括号就是 2n 长度)
		// 说明生成了一个合法的组合，直接加入结果集
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		// 铁律 1：只要左括号没用完，就可以一直放左括号
		if left < n {
			// 深入下一层，并且直接把 "(" 拼接到 path 里传下去
			backtrack(left+1, right, path+"(")
		}

		// 铁律 2：只要右括号比左括号少，就可以放右括号来闭合它
		if right < left {
			// 深入下一层，并且直接把 ")" 拼接到 path 里传下去
			backtrack(left, right+1, path+")")
		}
	}

	// 初始状态：用了 0 个左括号，0 个右括号，空字符串
	backtrack(0, 0, "")

	return res
}
