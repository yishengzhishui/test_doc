package leetcode

// 利用hash表，dp中key是当前石头的位置，value是跳几步可以到此石子的集合
// 跳到当前石子stone,所用的步数集合dp[stone]，遍历此集合下一步跳都有三个选择
// 判断stone+k是否在stonesSet中
// 最后一个位置对应的集合非空，那也就意味着我们可以到达终点
func canCrossV1(stones []int) bool {
	// 如果第二个石头的位置不是1，直接返回false
	if stones[1] != 1 {
		return false
	}

	// 动态规划表，dp[i][j]表示在石头位置i处，上一步跳跃了j步
	dp := make(map[int]map[int]bool)
	// 记录石头的位置，用于快速查找
	stonesSet := make(map[int]bool)

	// 初始化动态规划表和石头集合
	for _, stone := range stones {
		dp[stone] = make(map[int]bool)
		// 用于快速判断某个位置是否有石头。
		stonesSet[stone] = true
	}

	// 初始位置0，上一步跳跃0步是可行的
	dp[0][0] = true

	// 遍历石头数组，计算动态规划表
	for i := 0; i < len(stones)-1; i++ {
		// 当前石头位置
		stone := stones[i]
		// 遍历在当前石头位置上可以跳跃的步数
		// dp[stone] 存储在当前石头位置上可以跳跃的步数集合
		for jump := range dp[stone] {
			// 遍历当前步数的前一步、当前步数、以及后一步
			for k := jump - 1; k <= jump+1; k++ {
				// 如果步数合法且下一步的石头位置存在，标记在下一步位置上跳跃了k步
				if k > 0 && stonesSet[stone+k] {
					dp[stone+k][k] = true
				}
			}
		}
	}

	// 判断最后一个石头位置上是否存在可行的跳跃步数
	return len(dp[stones[len(stones)-1]]) > 0
}

// BFS + 记忆化搜索
// stack放当前石子位置和上一步到这儿的步数
// fail 放已经放过stack的石子位置和步数，即已经走过的不可再用
func canCross(stones []int) bool {
	// 如果第二个石头的位置不是1，直接返回false
	if stones[1] != 1 {
		return false
	}

	// 记录已经失败的石头位置和跳跃步数，防止重复计算
	fail := make(map[int]map[int]bool)
	// 记录石头的位置，用于快速查找
	stonesSet := make(map[int]bool)
	for _, stone := range stones {
		stonesSet[stone] = true
	}
	// 使用栈来进行深度优先搜索
	stack := [][]int{{0, 0}}

	// 遍历栈，进行深度优先搜索
	for len(stack) > 0 {
		// 弹出栈顶石头位置和跳跃步数
		stoneJump := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		stone, jump := stoneJump[0], stoneJump[1]
		// 遍历当前石头位置上可能的跳跃步数
		for k := jump - 1; k <= jump+1; k++ {
			// 计算下一步的石头位置
			s := stone + k
			// 如果跳跃步数合法，下一步的石头位置存在，并且之前没有失败记录
			if k > 0 && stonesSet[s] && !fail[s][k] {
				// 如果下一步的石头位置是最后一个石头，返回true
				if s == stones[len(stones)-1] {
					return true
				}
				// 将下一步的石头位置和跳跃步数入栈
				stack = append(stack, []int{s, k})
			}
		}
		// 记录当前石头位置和跳跃步数为失败状态
		if _, exists := fail[stone]; !exists {
			fail[stone] = make(map[int]bool)
		}
		fail[stone][jump] = true
	}

	// 搜索结束，未找到可行解，返回false
	return false
}
