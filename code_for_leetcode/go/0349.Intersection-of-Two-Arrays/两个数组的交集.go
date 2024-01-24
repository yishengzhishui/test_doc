package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	// 创建一个 map 用于存储 nums1 中的元素，以及结果数组 res
	m := map[int]bool{}
	var res []int

	// 第一次遍历 nums1，将每个元素存储到 map 中，键为元素值，值为 true
	for _, n := range nums1 {
		m[n] = true
	}

	// 第二次遍历 nums2，如果元素在 map 中存在，则删除该元素并将其添加到结果数组 res 中
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}

	// 返回结果数组 res，即为两个数组的交集
	return res
}
