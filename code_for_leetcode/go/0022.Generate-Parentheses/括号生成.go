package leetcode

// 递归 回溯算法+剪枝
// 左括号肯定是要比右括号先用完
// 有效的括号组合
// 肯定是左括号先开始
func generateParenthesis(n int) []string {
	// 定义结果切片
	var result []string

	// 定义递归辅助函数
	var DFS func(left, right int, cur string)
	DFS = func(left, right int, cur string) {
		// 当右括号数量等于 n 时，将当前组合添加到结果切片
		if right == n {
			result = append(result, cur)
			return
		}

		// 如果左括号数量小于 n，则可以添加一个左括号
		if left < n {
			DFS(left+1, right, cur+"(")
		}

		// 如果右括号数量小于左括号数量，则可以添加一个右括号
		if right < left {
			DFS(left, right+1, cur+")")
		}
	}

	// 初始调用递归函数
	DFS(0, 0, "")

	// 返回最终结果切片
	return result
}
