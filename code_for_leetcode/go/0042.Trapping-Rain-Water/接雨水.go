package leetcode

func trap(height []int) int {
	// faster and better 双指针
	// leftMax：左边的最大值，它是从左往右遍历找到的
	// rightMax：右边的最大值，它是从右往左遍历找到的
	// left：从左往右处理的当前下标
	// right：从右往左处理的当前下标
	// 如果leftMax<rightMax，那么它就知道自己能存多少水了。
	//无论右边将来会不会出现更大的right_max，都不影响这个结果

	size, result := len(height), 0
	if size < 3 {
		return result // 如果数组长度小于3，表示无法容纳雨水，直接返回0
	}
	//left 和 right 是分别从数组的左右两端向中间移动的指针
	left, right := 0, size-1  // 左右两个指针分别指向数组的头和尾
	leftMax, rightMax := 0, 0 // 左右两边的最大高度，初始化为0

	for left <= right {
		if leftMax < rightMax {
			// 如果左边的最大高度小于右边的最大高度
			// 计算左边当前位置能够容纳的雨水，并更新左边的最大高度
			result += getMax(0, leftMax-height[left])
			leftMax = getMax(leftMax, height[left])
			left++
		} else {
			// 如果右边的最大高度小于等于左边的最大高度
			// 计算右边当前位置能够容纳的雨水，并更新右边的最大高度
			result += getMax(0, rightMax-height[right])
			rightMax = getMax(rightMax, height[right])
			right--
		}
	}

	return result // 返回最终的雨水面积

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
