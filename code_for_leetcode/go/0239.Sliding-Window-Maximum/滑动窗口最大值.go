package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	// 双向队列
	// 当下标大于等于 k-1 时，result 开始传入数据，即当前的最大值
	// 队列没有必要维护窗口里的所有元素，只需要维护有可能成为窗口里最大值的元素就可以了，同时保证队列里的元素数值是由大到小的。
	var (
		queue  []int // 用于维护窗口内元素的下标
		result []int // 用于存储每个窗口的最大值
	)

	for i := 0; i < len(nums); i++ {
		// 保证队列单调递减
		// 如果当前元素比队尾元素大，说明当前元素可能成为后续窗口的最大值
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1] // 弹出队尾元素，保持单调递减
		}

		// 将当前下标加入队列
		queue = append(queue, i)

		// 如果当前下标减去队首下标大于等于 k，说明需要弹出队首元素
		if i-queue[0] >= k {
			queue = queue[1:] // 弹出队首元素
		}

		// 当下标大于等于 k-1 时，开始将当前最大值加入结果
		if i+1 >= k {
			result = append(result, nums[queue[0]]) // 将队首元素对应的值加入结果
		}
	}

	return result
}
