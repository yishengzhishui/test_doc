package leetcode

// 二分查找
// 因为其实整个数组是连续的
// 使用 // 向下取整比较方便
func searchMatrix(matrix [][]int, target int) bool {
	// 如果矩阵为空，直接返回 false
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 获取矩阵的行数和列数
	// m 是行数，n是列数
	m, n := len(matrix), len(matrix[0])

	// 初始化左右指针
	left, right := 0, m*n-1

	// 开始二分查找
	for left <= right {
		mid := (left + right) / 2
		// 计算当前位置在二维矩阵中的行和列
		row, col := mid/n, mid%n
		num := matrix[row][col]

		// 如果找到目标值，返回 true
		if num == target {
			return true
		}

		// 如果当前值大于目标值，更新右指针
		if num > target {
			right = mid - 1
		} else {
			// 如果当前值小于目标值，更新左指针
			left = mid + 1
		}
	}

	// 如果循环结束未找到目标值，返回 false
	return false
}
