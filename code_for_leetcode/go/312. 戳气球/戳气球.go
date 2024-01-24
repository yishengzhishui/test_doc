package leetcode

// dp[i][j] = x 表示，戳破气球 i 和气球 j 之间（开区间，不包括 i 和 j）的所有气球，可以获得的最高分数为 x。
// “开区间” 的意思是，我们只能戳爆 i 和 j 之间的气球，i 和 j 不要戳
// 如果最后一个戳破气球 k，dp[i][j] 的值应该为：
// dp[i][j] = getMax(dp[i][j], dp[i][k] + dp[k][j] + nums[i] * nums[j] * nums[k])
// 为了保证所有的依赖项被算出，需要注意i，j的遍历顺序：从下到上；从左到右
// maxCoins 函数计算戳气球游戏的最高得分。
func maxCoins(nums []int) int {
	// 在数组 nums 前后添加 1，表示两个虚拟气球，使得最终只剩下气球 0 和气球 len(nums)-1 两个气球。
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	size := len(nums)

	// 使用二维数组 dp 表示戳破气球 i 和气球 j 之间的所有气球，可以获得的最高分数。
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	// 通过三重循环遍历 i、j 和 k，更新 dp[i][j] 的值。
	//j在i后面，所以i从size-2开始
	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = getMax(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])
			}
		}
	}

	// 返回 dp[0][size-1]，表示从气球 0 到气球 len(nums)-1 可以获得的最高分数。
	return dp[0][size-1]
}

// getMax 函数用于返回两个整数中的较大值。
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
