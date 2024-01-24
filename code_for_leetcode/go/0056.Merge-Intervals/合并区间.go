package leetcode

// 排序+双指针

import "sort"

func merge(intervals [][]int) [][]int {
	// 先排序
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照左边界 先排序
	// 使用 sort.Slice 函数按照区间的起始值进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for _, e := range intervals {
		// 检查当前区间的起始值是否大于结果集中最后一个区间的结束值
		// 如果 当前区间 左边界大于 最后一个区间的右边界， 直接放入
		if e[0] > result[len(result)-1][1] {
			result = append(result, e)
		} else {
			// 合并区间
			// 因为intervals都是排序过的，以左边界排过序了
			result[len(result)-1][1] = getMax(result[len(result)-1][1], e[1])
		}
	}

	return result
}

// 辅助函数，返回两个数的最大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
