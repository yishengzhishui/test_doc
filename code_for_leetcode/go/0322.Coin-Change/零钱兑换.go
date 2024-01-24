package leetcode

import (
	"math"
	"sort"
)

// 类比上楼梯问题，
// dp[i]是凑齐价值为i需要的最少硬币数量
// 首先将dp[i]设置为一个不可能的数
// 想凑到amount，那么凑到amount-coins[i]时的硬币个数+1
// 可能的组合数类似 DP f[n] = f[n-1]+f[n-2]
func coinChange(coins []int, amount int) int {
	// 初始化一个切片 dp，用于存储凑齐不同价值的最少硬币数量
	dp := make([]int, amount+1)

	// 将 dp 中的所有元素初始化为一个不可能的数
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// 初始状态，凑齐价值为 0 的硬币数量为 0
	dp[0] = 0

	// 遍历硬币集合
	for _, coin := range coins {
		// 遍历凑齐不同价值的硬币数量
		for i := coin; i <= amount; i++ {
			// 更新最小硬币数量
			dp[i] = getMin(dp[i], dp[i-coin]+1)
		}
	}

	// 如果最终的最小硬币数量仍为不可能的数，则无法凑齐目标金额
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	// 返回凑齐目标金额所需的最少硬币数量
	return dp[amount]
}

// 辅助函数，返回两个整数中的较小值
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// DFS+剪枝
// 先丢大的硬币
func coinChangeV1(coins []int, amount int) int {
	// 将硬币集合按降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// 初始化最终结果
	res := math.MaxInt32

	// 定义深度优先搜索函数 DFS
	var DFS func(x, amount, count int)
	DFS = func(x, amount, count int) {
		// 如果 amount 为 0，则更新最终结果
		if amount == 0 {
			res = min(res, count)
			return
		}

		// 遍历硬币集合
		for i := x; i < len(coins); i++ {
			// 判断条件，确保 amount 大于当前硬币面值，并且 amount 没有超过当前最优解的上限
			if coins[i] <= amount && amount < coins[i]*(res-count) {
				DFS(i, amount-coins[i], count+1)
			}
		}
	}

	// 调用 DFS 函数开始深度优先搜索
	DFS(0, amount, 0)

	// 返回最终结果，如果仍然为初始值，则说明无法凑齐目标金额
	if res == math.MaxInt32 {
		return -1
	}

	return res
}
