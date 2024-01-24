package leetcode

//0011_盛最多水的容器
func maxArea(height []int) int {
	// 双指针 夹逼法
	// 移动 较小的边

	maxArea := 0                    // 用于存储最大的面积
	left, right := 0, len(height)-1 // 左右两个指针，分别指向数组的头和尾

	for left < right { // 当左指针小于右指针时，进行循环
		var (
			area      int // 存储当前计算的面积
			curHeight int // 存储当前高度
		)
		width := right - left // 计算宽度

		if height[left] < height[right] { // 如果左边的高度小于右边的高度
			curHeight = height[left] // 当前高度取左边的高度

			left++ // 移动左指针，缩小宽度
		} else {
			curHeight = height[right] // 当前高度取右边的高度
			right--                   // 移动右指针，缩小宽度
		}

		area = curHeight * width            // 计算当前的面积
		maxArea = getMaxArea(maxArea, area) // 更新最大面积
	}

	return maxArea // 返回最大面积
}

// max 函数，用于返回两个整数中的较大值
func getMaxArea(a, b int) int {
	if a > b {
		return a
	}
	return b
}
