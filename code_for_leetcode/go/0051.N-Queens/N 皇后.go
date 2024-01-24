package leetcode

import (
	"strings"
)

//回溯算法
// 同一斜线 x,y的特点是：x + y 的值相等， x-y的值也相等
// 所以保证queue的 x+y，x-y不能有相等的

func solveNQueens(n int) [][]string {
	// 定义结果切片
	var result [][]int

	// 定义递归辅助函数
	// cur: 当前解决方案的一行，记录每行皇后所在的列。
	// xyDiff: 记录斜线 "x-y" 的差值（左上到右下的对角线），用于检查是否存在冲突。
	// xySum: 记录斜线 "x+y" 的和值（左下到右上的对角线），用于检查是否存在冲突。
	var helper func(xyDiff, xySum []int, cur []int)
	helper = func(xyDiff, xySum []int, cur []int) {
		// 获取当前行数
		row := len(cur)

		// 如果已经排列完所有行，将当前解添加到结果切片
		if row == n {
			result = append(result, append([]int{}, cur...))
			return
		}

		// 遍历当前行的每一列(对于每一行遍历每一列)
		for col := 0; col < n; col++ {
			// 检查当前列、左对角线和右对角线是否有冲突
			if !contains(cur, col) && !contains(xySum, col+row) && !contains(xyDiff, col-row) {
				// 添加当前列到解中
				helper(append(xyDiff, col-row), append(xySum, col+row), append(cur, col))
			}
		}
	}

	// 初始化调用递归函数
	helper([]int{}, []int{}, []int{})

	//格式化结果切片为字符串切片
	var formattedResult [][]string
	for _, solution := range result {
		// formatted 用于存储格式化后的当前解决方案
		formatted := make([]string, n)

		// 遍历当前解决方案的每一列
		for i, col := range solution {
			// 生成当前列的字符串表示，"Q" 表示皇后，"." 表示空格
			formatted[i] = strings.Repeat(".", col) + "Q" + strings.Repeat(".", n-col-1)
		}

		// 将格式化后的当前解决方案添加到 formattedResult 中
		formattedResult = append(formattedResult, formatted)
	}

	// 返回最终结果
	return formattedResult
}

func contains(arr []int, target int) bool {
	// 检查切片中是否包含目标值
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}
