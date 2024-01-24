package leetcode

// dp[i] 走到当前阶梯的话费；等于本阶梯的花费+前面两个的较小者
// dp[i] = cost[i] + min(dp[i-1], dp[i-2])
// 最后比较dp[-1]和[-2]，因为到顶可以选择走一步或两步

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	size := len(cost)
	dp := make([]int, size)

	for i := 0; i < size; i++ {
		if i < 2 {
			dp[i] = cost[i]
		} else {
			dp[i] = cost[i] + getMin(dp[i-1], dp[i-2])
		}
	}
	return getMin(dp[size-1], dp[size-2])
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
