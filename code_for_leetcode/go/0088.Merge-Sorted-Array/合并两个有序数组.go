package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 使用指针 p，初始位置为 m+n
	for p := m + n; m > 0 && n > 0; p-- {
		// 比较 nums1[m-1] 和 nums2[n-1] 的大小
		if nums1[m-1] <= nums2[n-1] {
			// 如果 nums1[m-1] 小于等于 nums2[n-1]，将 nums2[n-1] 放到 nums1[p-1] 的位置
			nums1[p-1] = nums2[n-1]
			n-- // 将 nums2 中的指针 n 向前移动
		} else {
			// 如果 nums1[m-1] 大于 nums2[n-1]，将 nums1[m-1] 放到 nums1[p-1] 的位置
			nums1[p-1] = nums1[m-1]
			m-- // 将 nums1 中的指针 m 向前移动
		}
	}

	// 如果 nums2 中还有剩余元素，将其复制到 nums1 的前面
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// mergeSortedArrays 函数用于将两个有序数组合并到第一个数组中
func mergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	// 初始化两个数组的索引和合并后数组的索引
	i, j := m-1, n-1
	mergedIndex := m + n - 1

	// 循环比较两个数组的元素，从后向前合并
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			// 如果 nums1 的当前元素大于 nums2 的当前元素，将 nums1 的元素复制到合并后数组的位置
			nums1[mergedIndex] = nums1[i]
			i--
		} else {
			// 如果 nums2 的当前元素大于等于 nums1 的当前元素，将 nums2 的元素复制到合并后数组的位置
			nums1[mergedIndex] = nums2[j]
			j--
		}
		// 更新合并后数组的索引
		mergedIndex--
	}

	//如果nums2有剩余，就是存在有元素比nums1中所有元素都小
	//处理 nums2 数组剩余的元素，如果 j>=0，则将 nums2 的剩余元素复制到 nums1 的前面
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}
