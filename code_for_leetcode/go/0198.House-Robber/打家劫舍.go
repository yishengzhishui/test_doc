package leetcode

// 解法一 DP
// DP: 1、重复子问题；2、状态的定义；3、DP方程

// a[i][0,1]: 0:i偷， 1:i不偷
// a[i][0] = getMax(a[i-1][0], a[i-1][1])
// a[i][1] = a[i-1][0] + nums[i]

func robV1(nums []int) int {
	// 判断输入切片是否为空
	if len(nums) == 0 {
		return 0
	}

	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	dp := make([][2]int, size)

	// 初始化
	dp[0][1] = nums[0]

	// 循环计算状态数组 dp 中的值
	for i := 1; i < size; i++ {
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return getMax(dp[size-1][0], dp[size-1][1])
}

// getMax 函数返回两个整数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化：不增加新的维度, 记录每一步的最优结果
// a[i]: 0..1 能够偷到 MAX Value，结果返回 max(a)
// a[i] = max(a[i-1]+0, a[i-2]+nums[i]) 偷i-1或是偷i和i-2
func rob(nums []int) int {
	// 判断输入切片是否为空
	if len(nums) == 0 {
		return 0
	}

	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size)

	// 初始化
	dp[0] = nums[0]
	dp[1] = getMax(nums[0], nums[1])

	// 循环计算状态数组 dp 中的值
	for i := 2; i < size; i++ {
		dp[i] = getMax(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[size-1]
}
