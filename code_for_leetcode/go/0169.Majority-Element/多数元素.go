package leetcode

// Boyer-Moore 投票算法
// 投票法是遇到相同的则票数 + 1，遇到不同的则票数 - 1。
// 因此“多数元素”的个数 - 其余元素的个数总和 的结果 肯定 >= 1。
// 多数元素 个数是大于 len(nums)/2
func majorityElementV1(nums []int) int {
	candidate, count := 0, 0

	// 遍历整个数组
	for i := 0; i < len(nums); i++ {
		// 如果计数为零，更新候选值
		if count == 0 {
			candidate = nums[i]
		}

		// 如果当前元素与候选值相等，增加计数，否则减少计数
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate // 返回最终的候选值
}

// 使用hash 两次遍历
func majorityElementWithMap(nums []int) int {
	count := make(map[int]int)

	// 遍历整个数组，统计每个元素的出现次数
	for _, num := range nums {
		count[num]++
	}

	maxCount, majority := 0, 0

	// 遍历统计结果，找到出现次数最多的元素
	for key, value := range count {
		if value > maxCount {
			maxCount = value
			majority = key
		}
	}

	return majority // 返回出现次数最多的元素
}
