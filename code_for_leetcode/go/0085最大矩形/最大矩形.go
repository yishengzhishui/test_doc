package leetcode

// 与84类似 需要进行变换
// 将matrix的每一行看作地面，这一列上面的数值加起来就是84题目中矩形的高度
// 如果当前行的列是0，则高度就降为0， 否则高度是可以累加的

func maximalRectangle(matrix [][]byte) int {
	// 检查矩阵是否为空
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	// 获取矩阵的行数和列数
	m, n := len(matrix), len(matrix[0])

	// 初始化矩形的高度数组 heights，其长度比列数多一列
	heights := make([]int, n+1)

	// 初始化最大矩形面积为 0
	result := 0

	// 遍历矩阵的每一行
	for i := 0; i < m; i++ {
		// 初始化栈，初始时放入 -1 作为边界标记
		stack := []int{-1}

		// 更新矩形的高度数组 heights
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		// 遍历 heights 数组，计算每一列的最大矩形面积
		for x := 0; x < n+1; x++ {
			// 当前高度小于栈顶元素对应的高度时，表示找到了右边界
			for len(stack) > 1 && heights[x] < heights[stack[len(stack)-1]] {
				// 出栈并计算矩形面积
				height := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				width := x - stack[len(stack)-1] - 1
				// 更新结果为当前矩形面积和之前结果中的较大值
				result = getMAx(result, height*width)
			}
			// 当前列的索引入栈
			stack = append(stack, x)
		}
	}

	// 返回最终的最大矩形面积
	return result
}

// getMAx 函数用于返回两个整数中的较大值
func getMAx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
