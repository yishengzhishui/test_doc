package leetcode

// 给的元素会重复，需要去重

// permuteUnique 函数接受一个整数切片 nums，返回一个包含 nums 中所有数字的全排列的切片的切片。
func permuteUnique(nums []int) [][]int {
	// 初始化结果切片
	var result [][]int
	// 获取 nums 的长度
	size := len(nums)

	// 定义嵌套函数 helper，用于递归生成全排列
	var helper func(cur []int, numsMap map[int]int)
	helper = func(cur []int, numsMap map[int]int) {
		// 如果当前组合长度等于 nums 的长度，将其加入结果切片
		if len(cur) == size {
			result = append(result, append([]int{}, cur...))
			return
		}

		// 遍历当前可选的数字
		for num := range numsMap {
			// 如果该数字仍有剩余可用次数
			if numsMap[num] > 0 {
				// 减少该数字的可用次数
				numsMap[num]--
				// 将当前数字添加到新的组合中
				newCur := append([]int{}, cur...)
				newCur = append(newCur, num)
				// 递归调用 helper，生成下一层的排列
				helper(newCur, numsMap)
				// 恢复该数字的可用次数（如果不恢复，for循环就继续不下去了）
				numsMap[num]++
			}
		}
	}

	// 使用 Counter 计算每个数字的出现次数
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}

	// 调用 helper 函数，开始递归生成全排列
	helper([]int{}, numsMap)

	// 返回最终的结果切片
	return result
}
