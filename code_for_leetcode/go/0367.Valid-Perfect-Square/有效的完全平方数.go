package leetcode

// 二分查找
// 数的平方在坐标轴的右侧是单调递增的；
// 并且结果是存在上下边界的，结果肯定在1~x之间
// 和 69 x的平方根 类似

func isPerfectSquare(num int) bool {
	// 初始化左右边界
	left, right := 1, num

	// 循环条件：左边界小于等于右边界
	for left <= right {
		// 计算中间值
		mid := (left + right) / 2

		// 计算中间值的平方
		x := mid * mid
		if x == num {
			return true
		}

		// 如果中间值的平方大于目标值 m，说明解在左半部分
		if x > num {
			right = mid - 1
		} else {
			// 否则，解在右半部分或者中间值即为答案
			left = mid + 1
		}
	}

	return false
}
