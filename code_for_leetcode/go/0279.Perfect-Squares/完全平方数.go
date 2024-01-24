package leetcode

// 和凑硬币类似
// dp[i] = min(dp[i], dp[i-$nums[j]]+1)
// 先找的小于n的 最大的sqrt, 然后按照凑硬币的方法

func numSquares(n int) int {
	res := n

	// 初始化动态规划数组 dp，长度为 n+1
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0

	// 使用牛顿法求解平方根
	for res*res > n {
		res = (res + n/res) / 2
	}

	// 动态规划计算最小的平方数数量
	for e := 1; e <= res; e++ {
		for i := e * e; i <= n; i++ {
			dp[i] = getMin(dp[i], dp[i-e*e]+1)
		}
	}

	return dp[n]
}

// 辅助函数，返回两个整数中的最小值
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
