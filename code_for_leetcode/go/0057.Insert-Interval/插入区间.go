package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	left, right := newInterval[0], newInterval[1]

	for index, value := range intervals {
		//如果当前 value 的右边界小于 new的左边界，不需要合并区间
		if value[1] < left {
			result = append(result, value)
		} else if value[0] > right {
			// 如果当前 value 的左边界大于 new的右边界，不需要合并区间，new也可以放入了
			result = append(result, []int{left, right})
			// new放入了，后面也可以直接放进结果
			return append(result, intervals[index:]...)
		} else {
			// 合并边界
			left = getMin(left, value[0])
			right = getMax(right, value[1])
		}
	}

	// 上面遍历结束，新的区间放在尾部
	result = append(result, []int{left, right})
	return result
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
