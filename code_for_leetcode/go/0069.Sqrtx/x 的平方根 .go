package leetcode

// 解法一 
// 二分查找
// 数的平方在坐标轴的右侧是单调递增的；
// 并且结果是存在上下边界的，结果肯定在1~x之间
// 当right<left时，结束循环，left和right只相差1了，返回right就是取小的那个

func mySqrt(x int) int {
	// 初始化左右边界
	left, right := 1, x

	// 循环条件：左边界小于等于右边界
	for left <= right {
		// 计算中间值
		mid := (left + right) / 2

		// 计算中间值的平方
		num := mid * mid

		// 如果中间值的平方大于目标值 x，说明解在左半部分
		if num > x {
			right = mid - 1
		} else {
			// 否则，解在右半部分或者中间值即为答案
			left = mid + 1
		}
	}

	// 返回结果，注意这里返回的是右边界 right，因为题目要求返回整数部分
	return right
}

// 解法二 
// 牛顿迭代法
// 参考1：https://leetcode-cn.com/problems/sqrtx/solution/er-fen-cha-zhao-niu-dun-fa-python-dai-ma-by-liweiw/
// 参考2：https://leetcode-cn.com/problems/sqrtx/solution/niu-dun-die-dai-fa-by-loafer/
// 牛顿法的思想：
// 在迭代过程中，以直线代替曲线，用一阶泰勒展式（即在当前点的切线）代替原曲线，求直线与 xx 轴的交点，重复这个过程直到收敛。
// 迭代的公式就是 (x+a/x)/2，不断逼近这个结果
// 牛顿法得到的是平方根的浮点型精确值 需要去掉小数部分
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
