package leetcode

// DP
// dp[i][j] = getMin(dp[i+1][j],dp[i][j+1])+grid[i][j]
// basic dp
func minPathSum(grid [][]int) int {
	// 检查输入是否为空
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 获取行和列的数量
	m, n := len(grid), len(grid[0])
	// 创建二维数组 dp 用于记录最小路径和
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 其中 dp[i][j] 表示从位置 (i, j) 到右下角的最小路径和。
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 如果当前位置是右下角的位置
			if i == m-1 && j == n-1 {
				dp[i][j] = grid[i][j]
			} else if i == m-1 {
				// 如果当前位置在最后一行，只能往右移动
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == n-1 {
				// 如果当前位置在最后一列，只能往下移动
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else {
				// 如果当前位置不在最后一行或最后一列，可以选择往右或往下移动，选择最小路径和
				dp[i][j] = grid[i][j] + getMin(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 原地，不用额外空间
// dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
func minPathSumV1(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 更新第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// 更新第一行
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 更新其余位置
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += getMin(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
