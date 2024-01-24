package leetcode

// DP
// f[i][j] = （f[i - 1][j] + f[i][j - 1]）*（1-障碍物）

func uniquePathsWithObstaclesV1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 初始化动态规划数组 dp，大小为 m x n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起点，如果起点有障碍物，dp[0][0] 为 0，否则为 1
	dp[0][0] = 1 - obstacleGrid[0][0]

	// 初始化第一列，如果有障碍物，则后面的 dp[i][0] 都为 0，否则 dp[i][0] 等于上一行的 dp[i-1][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] * (1 - obstacleGrid[i][0])
	}

	// 初始化第一行，如果有障碍物，则后面的 dp[0][j] 都为 0，否则 dp[0][j] 等于前一列的 dp[0][j-1]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] * (1 - obstacleGrid[0][j])
	}

	// 动态规划递推
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果当前位置有障碍物，dp[i][j] 置为 0
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				// 否则，dp[i][j] 为左边和上边路径的和
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// 返回右下角的值，即为从起点到终点的不同路径数
	return dp[m-1][n-1]
}

// 优化表达
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 初始化动态规划数组 dp，大小为 m x n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起点，如果起点有障碍物，dp[0][0] 为 0，否则为 1
	dp[0][0] = 1 - obstacleGrid[0][0]

	// 动态规划递推
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前位置有障碍物，dp[i][j] 置为 0
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				// 根据动态规划递推式，更新当前位置的路径数
				if i > 0 {
					dp[i][j] += dp[i-1][j] * (1 - obstacleGrid[i][j])
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1] * (1 - obstacleGrid[i][j])
				}
			}
		}
	}

	// 返回右下角的值，即为从起点到终点的不同路径数
	return dp[m-1][n-1]
}
