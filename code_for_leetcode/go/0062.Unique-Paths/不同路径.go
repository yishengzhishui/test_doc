package leetcode

// dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePaths(m, n int) int {
	// 创建一个二维数组 dp，大小为 m 行 n 列，用于存储到达每个位置的路径数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 将起始位置 (0, 0) 的路径数设置为 1，因为从左上角到左上角只有一种路径
	dp[0][0] = 1

	// 遍历矩阵的每个位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前位置不在第一行，可以从上方到达，将上方位置 (i-1, j) 的路径数加到当前位置的路径数
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			// 如果当前位置不在第一列，可以从左侧到达，将左侧位置 (i, j-1) 的路径数加到当前位置的路径数
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	// dp[m-1][n-1] 存储的是从左上角到右下角的不同路径数，函数返回这个值
	return dp[m-1][n-1]
}
