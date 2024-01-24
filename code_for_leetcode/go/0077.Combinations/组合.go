package leetcode

// 回溯算法
func combineV1(n, k int) [][]int {
	// 定义嵌套函数 helper，用于递归生成组合
	var helper func(cur []int, start int)

	// 初始化结果切片
	var result [][]int

	// 实现 helper 函数
	// cur是当前组合，start是当前数字起点
	helper = func(cur []int, start int) {
		// 如果当前组合长度等于 k，将其加入结果切片
		if len(cur) == k {
			//append([]int{}, cur...)为了创建 cur 切片的副本
			// 防止切片共享底层数组
			result = append(result, append([]int{}, cur...))
			return
		}

		// 遍历可能的数字，从 start 到 n
		for i := start; i <= n; i++ {
			// 递归调用 helper，生成下一层的组合
			helper(append(cur, i), i+1)
		}
	}

	// 调用 helper 函数，开始递归生成组合
	helper([]int{}, 1)

	// 返回最终的结果切片
	return result
}

// 回溯算法 + 减枝
func combine(n, k int) [][]int {
	// 定义嵌套函数 helper，用于递归生成组合
	var helper func(cur []int, start int)

	// 初始化结果切片
	var result [][]int

	// 实现 helper 函数
	// cur是当前组合，start是当前数字起点
	helper = func(cur []int, start int) {
		// 如果当前组合长度等于 k，将其加入结果切片
		if len(cur) == k {
			//append([]int{}, cur...)为了创建 cur 切片的副本
			// 防止切片共享底层数组
			result = append(result, append([]int{}, cur...))
			return
		}

		// 遍历可能的数字，从 start 到 n
		for i := start; i <= n; i++ {
			// 剪枝：剩余的不满足要求了
			// 举个例子，n = 4，k = 3， 目前已经选取的元素为0，
			//n - (k - 0) + 1 即 4 - ( 3 - 0) + 1 = 2。
			if n-i+1 < k-len(cur) {
				break
			}
			// 递归调用 helper，生成下一层的组合
			helper(append(cur, i), i+1)
		}
	}

	// 调用 helper 函数，开始递归生成组合
	helper([]int{}, 1)

	// 返回最终的结果切片
	return result
}
