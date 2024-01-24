package leetcode

func isValid(s string) bool {
	stack := []rune{'?'}                               // 初始化栈，包含一个不相关的值 '?'
	dic := map[rune]rune{'(': ')', '[': ']', '{': '}'} // 映射左右括号关系

	// 遍历输入字符串中的每个字符
	for _, char := range s {
		// 如果当前字符是左括号，将对应的右括号入栈
		if left, found := dic[char]; found {
			stack = append(stack, left)
		} else {
			// 检查栈顶元素是否与当前右括号匹配
			// 如果匹配，说明找到了一对匹配的括号，将栈顶元素出栈
			if char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// 检查栈的大小是否为 1，如果是，说明所有的左括号都有相应的右括号，返回 true；否则，返回 false
	return len(stack) == 1
}
