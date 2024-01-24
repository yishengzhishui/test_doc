package leetcode

import "sort"

func sortedSquaresV1(nums []int) []int {
	// 获取数组长度
	n := len(nums)
	// 创建一个与原数组长度相同的结果数组
	ans := make([]int, n)
	// 使用双指针，分别从数组的两端开始
	i, j := 0, n-1

	// 从结果数组的最后一个位置开始往前填充
	for pos := n - 1; pos >= 0; pos-- {
		// 获取左指针对应位置的平方值 v
		v := nums[i] * nums[i]
		// 获取右指针对应位置的平方值 w
		w := nums[j] * nums[j]

		// 比较平方值 v 和 w，将较大的值存入结果数组
		if v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}

	// 返回最终的结果数组
	return ans
}

// 解法一
func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, j, pos := 0, len(A)-1, len(ans)-1; i <= j; pos-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[pos] = A[i] * A[i]
			i++
		} else {
			ans[pos] = A[j] * A[j]
			j--
		}
	}
	return ans
}

// 解法二
func sortedSquares1(A []int) []int {
	for i, value := range A {
		A[i] = value * value
	}
	sort.Ints(A)
	return A
}
