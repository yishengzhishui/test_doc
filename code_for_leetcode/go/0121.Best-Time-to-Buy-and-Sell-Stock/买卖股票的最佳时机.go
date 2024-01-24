package leetcode

// 解法一 模拟 DP
// 买卖股票问题 套路公式  DP方法,穷举框架，列出所有的可能性的结果
//
// for 状态1 in 状态1的所有值：
//     for 状态2 in 状态2的所有值：
//         for .....:
//             dp[状态1][状态2][...] = 选择

// dp[i][k][o], i是天数，k是已经操作的次数，o=[0,1],0是手上无股票，1是有
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
// dp[i][0][1] 不存在 没有交易下是不可能持有股票的，不存在
// dp[i][0][0] = 0

//  k = 1 可以省略
func maxProfit(prices []int) int {
	// 如果价格切片为空，直接返回 0
	if len(prices) == 0 {
		return 0
	}

	// 获取价格切片的长度
	size := len(prices)

	// 初始化动态规划数组 dp
	dp := make([][2]int, size)

	// 初始化第一天的情况
	dp[0][1] = -prices[0]

	// 开始遍历价格切片
	for i := 1; i < size; i++ {
		// 计算当前天不持有股票和持有股票的最大利润
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = getMax(dp[i-1][1], -prices[i])
	}

	// 返回最后一天不持有股票的最大利润
	return dp[size-1][0]
}

// 辅助函数，返回两个整数中的较大值
func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
