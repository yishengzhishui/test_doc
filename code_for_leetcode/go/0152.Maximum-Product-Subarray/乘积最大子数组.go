package leetcode

// DP 因为存在负数，负负得正，所以需要维护一个最小值
// dp[i][0]getMax, dp[i][1]getMin
// 如果nums[i] < 0；dp[i][0]得到结果前需要dp[i-1][0], dp[i-1][1]交换数值；dp[i-1][0]就变成负的了
// 因为dp中的数据是会随时改变的，需要维护一个最大乘积result

func maxProduct(nums []int) int {
	// 空间优化
	if len(nums) == 0 {
		return 0
	}

	size := len(nums)
	result := nums[0]
	maxVal, minVal := result, result

	for i := 1; i < size; i++ {
		// 如果当前元素为负数，交换最大值和最小值
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}

		// 计算当前元素的最大值和最小值
		maxVal = getMax(maxVal*nums[i], nums[i])
		minVal = getMin(minVal*nums[i], nums[i])

		// 更新全局最大值
		result = getMax(result, maxVal)
	}

	return result
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
