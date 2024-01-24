package leetcode

// 第一家和最后一家，只能偷一家，其他的和198题目一致
// a[i]: 0..1 nums[i]是必偷时，能够偷到的最大值，结果返回 max(a)

// a[i] = max(a[i-1]+0, a[i-2]+nums[i])

func rob(nums []int) int {
	// 判断输入切片是否为空
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	size := len(nums)
	if size <= 2 {
		return getMax(nums[0], nums[1])
	}

	return getMax(helper(nums[1:]), helper(nums[:size-1]))
}

func helper(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	dp := make([]int, len(arr))

	// 初始化
	dp[0] = arr[0]
	dp[1] = getMax(arr[0], arr[1])

	// 循环计算状态数组 dp 中的值
	for i := 2; i < len(arr); i++ {
		dp[i] = getMax(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[len(arr)-1]
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
