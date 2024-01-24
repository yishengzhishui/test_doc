package leetcode

// stack, 类似于找左右边界，左边界在stack中
// 传入栈的是元素的下标，但是保证这些元素是单调递增的 height[stack[-1]是栈中最大的元素
// 遍历数组，如果heights[i] < heights[stack[-1]]，右边界找到;
// 一次遍历之后，需要考虑栈里的剩余元素的出栈，需要最后有一个最小的元素0，来帮助其他的元素出栈，所以heights.append(0)
// 为了减少边界条件判断，stack有初始值[-1]
func largestRectangleArea(heights []int) int {
	// 在输入数组末尾添加一个 0，以确保所有元素都能出栈
	heights = append(heights, 0)
	size := len(heights)

	// 初始化一个栈，初始时放入 -1 作为边界标记
	stack := []int{-1}
	result := 0

	for i := 0; i < size; i++ {
		// 当前元素小于栈顶元素时，出栈并计算矩形面积
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // 出栈
			width := i - stack[len(stack)-1] - 1
			result = getMax(result, height*width)
		}
		// 当前元素入栈
		stack = append(stack, i)
	}

	return result
}

// 辅助函数 getMax 用于返回两个整数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
