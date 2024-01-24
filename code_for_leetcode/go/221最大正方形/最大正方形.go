package leetcode

// dp 具体定义：dp[i + 1][j + 1] 表示 「以第 i 行、第 j 列为右下角的正方形的最大边长」
// if matrix[i - 1][j - 1] == '1':
//     dp[i][j] = getMin(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]) + 1
// 意思是：若某格子值为 1，则以此为右下角的正方形的最大边长为：上面的正方形、左面的正方形或左上的正方形中，最小的那个，再加上此格。
// 为了方便对于matrix[0][j]和matrix[i][0]判断，dp要多补全一行和一列 dp[height + 1][width + 1]



func maximalSquare(matrix [][]byte) int {
	// 如果矩阵为空，返回 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])

	// 初始化 dp 数组，多加一行一列，方便处理边界情况
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0 // 记录最大正方形的边长

	// 遍历矩阵
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果当前位置的值为 '1'
			if matrix[i-1][j-1] == '1' {
				// 计算当前位置的 dp 值，取左上、上、左三个相邻位置的最小值加一
				dp[i][j] = getMin(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				// 更新最大正方形的边长
				maxSide = getMax(maxSide, dp[i][j])
			}
		}
	}

	// 最大正方形的面积为边长的平方
	return maxSide * maxSide
}

// 辅助函数，返回两个整数中的较小值
func getMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// 辅助函数，返回两个整数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
