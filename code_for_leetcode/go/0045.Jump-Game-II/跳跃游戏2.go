package leetcode

// 贪心算法 从前往后
// 从头开始跳， 记录当前能到的最远距离 maxReach
// 如果 碰到i > 最远距离，则失败，若能一直跳下去就成功了
// 因为是依次遍历，这个需要保证max_reach是最大的，不能仅仅看i+nums[i]
// end指的是当前所能走到的最大的，当遍历到此位置是，step+1， end=max_reach

func jump(nums []int) int {
	// 初始化 maxReach 和 end，以及步数 step
	maxReach, end, step := 0, 0, 0

	// 循环遍历数组
	for i := 0; i < len(nums); i++ {
		// 如果当前位置超过了当前的可到达最远位置 end，那就是需要再跳一步
		// 则更新 end 为当前能够到达的最远位置 maxReach，并增加步数 step
		if i > end {
			end = maxReach
			step++
		}

		// 更新当前能够到达的最远位置 maxReach
		maxReach = getMax(maxReach, nums[i]+i)
	}

	// 返回最小步数 step
	return step
}

// getMax 返回两个数中的较大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
