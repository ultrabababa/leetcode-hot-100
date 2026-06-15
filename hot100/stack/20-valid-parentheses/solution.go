package main

func isValid(s string) bool {
	n := len(s)

	if n%2 != 0 {
		return false
	}

	stack := make([]int32, 0, n)
	for _, ch := range s {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
