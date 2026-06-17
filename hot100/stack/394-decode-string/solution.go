// hot100/stack/394-decode-string/solution.go

package main

import "strings"

// 用于记录在字符串中遇到[时需要保存的上下文状态
type state struct {
	k      int
	prefix string // 已生成的字符串
}

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var stack []state
	var curK int
	var curStr []byte
	// abc[4[a]3[b5[c]]]
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch >= '0' && ch <= '9':
			curK = curK*10 + int(ch-'0')
		case ch == '[':
			// 遇到左括号要保存当前状态并入栈
			stack = append(stack, state{
				k:      curK,
				prefix: string(curStr),
			})
			// 然后清空当前状态，准备记录括号里的内容
			curK = 0
			curStr = []byte{}
		case ch == ']':
			// 遇到右括号先弹出栈内保存的上一段状态信息
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			repeatedStr := strings.Repeat(string(curStr), top.k)
			curStr = []byte(top.prefix + repeatedStr)
		default:
			// 遇到普通字母
			curStr = append(curStr, ch)
		}
	}
	return string(curStr)
}
