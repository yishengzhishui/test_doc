package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	// 处理空字符串的情况
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	m, n := len(text1), len(text2)

	// 将字符串前面插入空格，使得索引从1开始
	text1 = " " + text1
	text2 = " " + text2

	// 初始化动态规划数组（默认都是0）
	//  dp[i][j] 表示 text1 的前 i 个字符和 text2 的前 j 个字符的最长公共子序列的长度。
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// 动态规划计算最长公共子序列长度
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 相同，找到了公共元素 dp+1
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				// 不相同，则比较取最大值
				dp[i][j] = getMax(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 返回最长公共子序列的长度
	return dp[m][n]
}

// 辅助函数，返回两个整数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
