package leetcode

// 贪心算法： 在每一步选择当前能够跳到的最远位置，从而判断是否能够跳到数组的起始位置。
// 从后向前
// 最远可以到达的位置 >= reach时， reach可以向前移动
// 当reach==0时，即从数组头部可以到达最后一个元素
func canJumpV1(nums []int) bool {
	// 初始化 reach 变量，表示当前能够到达的最远位置
	reach := len(nums) - 1

	// 从右向左遍历数组
	for i := reach; i >= 0; i-- {
		// 如果当前位置可以跳到最远位置，更新 reach
		if nums[i]+i >= reach {
			reach = i
		}
	}

	// 如果最终 reach 等于 0，说明能够到达起始位置
	return reach == 0
}

// 贪心策略，从头开始跳，记录当前能够到达的最远距离。
func canJump(nums []int) bool {
	// 从头开始跳，记录当前能到的最远距离
	// 如果碰到 i > 能到的最远距离，则失败，若能一直跳下去就成功了
	// 因为是依次遍历，这个需要保证 maxReach 是最大的，不能仅仅看 i+nums[i]
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = getMax(maxReach, nums[i]+i)
	}

	// 能够一直跳下去就成功了
	return true
}

// getMax 返回两个数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
