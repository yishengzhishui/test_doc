package leetcode

// 类比上楼梯问题，
// 注意：这里求的组合数，而不是排列数
// 我们不关心硬币使用的顺序，而是硬币有没有被用到
//动规五步曲来分析如下：
//1.确定dp数组以及下标的含义
//dp[j]：凑成总金额j的货币组合数为dp[j]
//2.确定递推公式
//dp[j] 就是所有的dp[j - coins[i]]（考虑coins[i]的情况）相加。
//所以递推公式：dp[j] += dp[j - coins[i]];
func change(amount int, coins []int) int {
	// 初始化动态规划数组 dp，长度为 amount+1
	dp := make([]int, amount+1)
	dp[0] = 1 // 初始情况，表示凑成金额为 0 的方法数为 1

	// 遍历硬币数组
	for _, coin := range coins {
		// 从当前硬币的面值开始，遍历到目标金额
		for i := coin; i <= amount; i++ {
			// 更新 dp[i]，表示凑成金额 i 的方法数
			// dp[i-coin] 表示在使用当前硬币前，已经凑到金额 i-coin 的方法数。
			dp[i] += dp[i-coin]
		}
	}

	// 返回凑成目标金额的方法数
	return dp[amount]
}
