package leetcode

import "strings"

// 使用stack
// 如果当前字符为 (，将其入栈；如果为 )，表示匹配了一个括号对，将栈顶的 ( 出栈。
func removeOuterParentheses(S string) string {
	// 使用 strings.Builder 作为结果字符串的构建器
	var result strings.Builder
	// 使用切片作为栈，用于记录括号的匹配状态
	stack := []byte{}
	// 记录当前有效括号序列的起始位置
	j := 0

	// 遍历输入字符串 S 中的每个字符
	for i := 0; i < len(S); i++ {
		// 如果当前字符为 '('，将其入栈
		if S[i] == '(' {
			stack = append(stack, S[i])
		} else if S[i] == ')' {
			// 如果当前字符为 ')'，表示匹配了一个括号对，将栈顶的 '(' 出栈
			stack = stack[:len(stack)-1]
		}

		// 如果栈为空，表示当前括号序列匹配完毕
		if len(stack) == 0 {
			// 提取有效括号序列（去除最外层括号），追加到结果字符串中
			result.WriteString(S[j+1 : i])
			// 更新 j 为下一个有效括号序列的起始位置
			j = i + 1
		}
	}

	// 返回最终结果字符串
	return result.String()
}

func RemoveOuterParentheses(S string) string {
	// 使用 strings.Builder 作为结果字符串的构建器
	var result strings.Builder
	// 用于记录当前括号序列的嵌套深度
	i := 0

	// 遍历输入字符串 S 中的每个字符
	for _, c := range S {
		// i>1 外层的那个已经遍历过了
		// 如果当前字符为 '(' 且嵌套深度大于 0，则将其追加到结果字符串中
		if c == '(' && i > 0 {
			result.WriteRune(c)
		}

		// 如果当前字符为 ')' 且嵌套深度大于 1，则将其追加到结果字符串中
		// i>1 至少还剩一个没有遍历（就是外层的那个）
		if c == ')' && i > 1 {
			result.WriteRune(c)
		}

		// 更新嵌套深度
		if c == '(' {
			i++
		} else {
			i--
		}
	}

	// 返回最终结果字符串
	return result.String()
}
