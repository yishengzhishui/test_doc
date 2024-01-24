package leetcode

func nthUglyNumber(n int) int {
	// 初始化三个指针 i2，i3，i5
	i2, i3, i5 := 0, 0, 0
	// 初始化数组 nums
	nums := []int{1}

	// 循环计算所有丑数
	for i := 1; i < n; i++ {
		// 在 nums[i2] * 2，nums[i3] * 3 和 nums[i5] * 5 选出最小的数字添加到数组 nums 中
		cur := getMin(nums[i2]*2, getMin(nums[i3]*3, nums[i5]*5))

		// 将该数字对应的因子指针向前移动一步
		if cur == nums[i2]*2 {
			i2++
		}
		if cur == nums[i3]*3 {
			i3++
		}
		if cur == nums[i5]*5 {
			i5++
		}

		// 将当前丑数添加到数组 nums 中
		nums = append(nums, cur)
	}

	// 返回第 n 个丑数
	return nums[n-1]
}

// 辅助函数，返回两个整数中的最小值
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
