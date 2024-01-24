package leetcode

import (
	"sort"
)

// topKFrequent 函数，返回出现频率前 k 高的元素
func topKFrequent(nums []int, k int) []int {
	// 使用 map 统计每个元素的出现频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 将 map 中的键值对转为切片，方便排序
	var freqList [][]int
	for num, freq := range freqMap {
		freqList = append(freqList, []int{num, freq})
	}

	// 对切片按照频率降序排序
	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i][1] > freqList[j][1]
	})

	// 取出前 k 个元素的键值
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, freqList[i][0])
	}

	return result
}
