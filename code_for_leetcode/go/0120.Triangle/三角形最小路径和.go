package leetcode

// 自底向上
// dp[i][j] 到当前位置的路径和
// dp[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
// 减少空间消耗可以triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])； 最后返回triangle[0][0]

// DP two-dimensional array
// 从倒数第二层开始遍历

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// 从倒数第二行开始逐层向上更新
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// 当前位置的值加上下一层相邻位置的最小值
			triangle[i][j] += getMin(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	// 最终结果存储在三角形的顶部
	return triangle[0][0]
}

func getMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
