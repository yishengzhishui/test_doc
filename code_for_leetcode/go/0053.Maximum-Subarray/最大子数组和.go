package leetcode

// dp[i] 是到第i个为止，连续子数组最大的和,最后返回dp中的最大值
// dp[i] = max(dp[i-1]+nums[i], nums[i])，
// num[i]是要加入到前面的dp,还是单独，取决于dp[i-1]+nums[i]是否对自己有增益，
// 即dp[i-1]+nums[i], nums[i]谁比较大

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	size := len(nums)
	dp := make([]int, size)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < size; i++ {
		dp[i] = getMax(dp[i-1]+nums[i], nums[i])
		result = getMax(result, dp[i])
	}
	return result

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//优化 cur shi
func maxSubArrayV1(nums []int) int {
	// 空间优化
	if len(nums) == 0 {
		return 0
	}

	size := len(nums)
	cur, result := nums[0], nums[0]

	for i := 1; i < size; i++ {
		cur = getMax(cur+nums[i], nums[i]) // 当前元素的最大子数组和，考虑是否继续扩展前面的子数组
		result = getMax(result, cur)       // 更新全局最大子数组和
	}

	return result
}
