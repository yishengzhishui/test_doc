package leetcode

import "sort"

func intersectV1(nums1 []int, nums2 []int) []int {
	// 使用 map 作为哈希表，用于记录 nums1 中每个元素的出现次数
	dic := make(map[int]int)
	// 用于存储交集元素的切片
	var res []int

	// 遍历 nums1，记录每个元素的出现次数
	for _, num := range nums1 {
		dic[num]++
	}

	// 遍历 nums2，寻找与 nums1 中相同元素的交集
	for _, num := range nums2 {
		// 如果 num 在 dic 中存在且次数大于 0，将其添加到交集中，并将次数减 1
		if count, ok := dic[num]; ok && count > 0 {
			res = append(res, num)
			dic[num]--
		}
	}

	// 返回最终的交集切片
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	// 对两个数组进行升序排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 检查数组是否为空
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	// 使用两个指针 i 和 j 分别遍历 nums1 和 nums2
	i, j := 0, 0
	// 用于存储交集元素的切片
	var res []int

	// 当 i 和 j 均在数组范围内时进行比较
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			// 如果当前元素相等，将其添加到交集中，并同时移动指针 i 和 j
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			// 如果 nums1[i] 小于 nums2[j]，移动指针 i
			i++
		} else {
			// 如果 nums1[i] 大于 nums2[j]，移动指针 j
			j++
		}
	}

	// 返回最终的交集切片
	return res
}
