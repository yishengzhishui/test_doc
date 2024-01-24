package leetcode

// removeElement 是移除指定值的函数
func removeElementV1(nums []int, val int) int {
	// 检查切片是否为空
	if len(nums) == 0 {
		return 0
	}

	// 初始化指针 j，用于指向新的位置
	j := 0

	// 遍历整个切片
	for i := 0; i < len(nums); i++ {
		// 如果当前元素不等于指定值 val
		if nums[i] != val {
			// 如果 i 不等于 j，表示当前元素需要移动到新的位置
			if i != j {
				// 交换元素，将不等于 val 的元素移动到前面
				nums[i], nums[j] = nums[j], nums[i]
			}
			// 指针 j 向后移动
			j++
		}
	}

	// 返回新的切片长度
	return j
}

func removeElement(nums []int, val int) int {
	// 检查切片是否为空
	if len(nums) == 0 {
		return 0
	}

	// 初始化指针 j，用于指向新的位置
	j := 0

	// 遍历整个切片
	for _, value := range nums {
		if val != value {
			nums[j] = value
			j++
		}
	}

	// 返回新的切片长度
	return j
}
