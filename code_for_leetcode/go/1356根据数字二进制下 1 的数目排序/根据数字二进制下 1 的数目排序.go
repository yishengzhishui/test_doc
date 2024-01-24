package leetcode

import (
	"math/bits"
	"sort"
)

// 位运算+hash
// 将二进制1的个数相同的元素放在同一个key下面
// 对dic的结果进行排序
func sortByBitsV1(arr []int) []int {
	// 创建 map 用于存储不同比特数的数字列表
	dic := make(map[int][]int)

	// 计算数字 x 的比特数
	bitCount := func(x int) int {
		count := 0
		for x > 0 {
			x &= x - 1
			count++
		}
		return count
	}

	// 将数字按比特数分类存储到 map 中
	for _, num := range arr {
		c := bitCount(num)
		//  bits.OnesCount 函数来计算数字的比特数，这是 Go 语言提供的标准库函数
		//c := bits.OnesCount(uint(num))
		dic[c] = append(dic[c], num)
	}

	// 对 map 中的比特数进行排序
	var keys []int
	for k := range dic {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 按照比特数和数字大小对结果进行排序
	var result []int
	for _, j := range keys {
		sortedNums := dic[j]
		sort.Ints(sortedNums)
		result = append(result, sortedNums...)
	}

	return result
}
//
func sortByBits(arr []int) []int {
	// 使用 Go 的排序函数，根据以下规则排序：
	// 1. 比特数较小的数字排在前面。
	// 2. 如果比特数相同，则数字较小的排在前面。
	sort.Slice(arr, func(i, j int) bool {
		// 计算 arr[i] 中的 1 的个数
		cx := bits.OnesCount(uint(arr[i]))
		// 计算 arr[j] 中的 1 的个数
		cy := bits.OnesCount(uint(arr[j]))

		// 使用比特数和数字大小进行比较
		// 如果比特数不同，返回比特数小的数字
		// 如果比特数相同，返回数字小的数字
		return cx < cy || (cx == cy && arr[i] < arr[j])
	})

	// 返回排序后的数组
	return arr
}
