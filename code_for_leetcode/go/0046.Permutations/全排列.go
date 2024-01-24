package leetcode

// permute 函数接受一个整数切片 nums，返回一个包含 nums 中所有数字的全排列的切片的切片。
func permute(nums []int) [][]int {
	// 初始化结果切片
	var result [][]int
	// 获取 nums 的长度
	size := len(nums)

	// 定义嵌套函数 helper，用于递归生成全排列
	var helper func(cur []int, arr []int)

	helper = func(cur []int, arr []int) {
		// 如果当前组合长度等于 nums 的长度，将其加入结果切片
		if len(cur) == size {
			result = append(result, append([]int{}, cur...))
			return
		}

		// 遍历当前可选的数字
		for i := 0; i < len(arr); i++ {
			// 复制当前组合
			newCur := append([]int{}, cur...)
			// 将当前数字添加到新的组合中
			newCur = append(newCur, arr[i])

			// 生成新的可选数字列表
			newArr := append([]int{}, arr[:i]...)
			newArr = append(newArr, arr[i+1:]...)

			// 递归调用 helper，生成下一层的排列
			helper(newCur, newArr)
		}
	}

	// 调用 helper 函数，开始递归生成全排列
	helper([]int{}, nums)

	// 返回最终的结果切片
	return result
}
