package leetcode

import (
	"math"
	"sort"
)

// 前缀和 + 排序
// 问题转化为 对于每一行 计算left到right的元素的和sums
// 随后对这一列中行元素和进行累加计算——前缀和，
//cur-k如果在array中，那么就返回max(result, cur - array[index])

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	result := math.MinInt64

	// 遍历所有可能的左边界
	for left := 0; left < n; left++ {
		sums := make([]int, m)

		// 遍历可能的右边界
		for right := left; right < n; right++ {
			// 计算每一行的和
			for i := 0; i < m; i++ {
				sums[i] += matrix[i][right]
			}

			// 利用有序数组进行查找和计算最大子矩阵和
			array := []int{0}
			cur := 0
			for _, num := range sums {
				cur += num

				// 在有序数组中二分查找满足条件的位置
				index := sort.Search(len(array), func(i int) bool { return array[i] >= cur-k })

				// 如果找到了满足条件的位置，更新结果
				if index < len(array) {
					result = max(result, cur-array[index])
				}

				// 将当前和插入有序数组
				array = insertSorted(array, cur)
			}
		}
	}
	return result
}

// 在有序数组中插入一个元素并保持有序性
// insertSorted 在有序数组中插入一个元素并保持有序性
func insertSorted(array []int, num int) []int {
	// 使用二分查找找到插入位置的索引
	index := sort.Search(len(array), func(i int) bool { return array[i] >= num })

	// 在数组末尾添加一个元素，确保有足够的空间
	array = append(array, 0)

	// 将插入位置之后的元素向后移动一位
	copy(array[index+1:], array[index:])

	// 在插入位置设置新的元素值
	array[index] = num

	// 返回更新后的有序数组
	return array
}
