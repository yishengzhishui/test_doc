package leetcode

// DFS+回溯算法
// 首先遍历数组，找到起点，终点；并且确认总共走要多少步（因为每个空格都要走一次且不可重复）
// 然后开始回溯
// 走到结束位置，如果step也为0，说明此路径成立

// 定义 uniquePathsIII 函数，输入是一个二维数组 grid，输出是符合条件的路径数量
func uniquePathsIII(grid [][]int) int {
	// 获取二维数组的行数和列数
	m, n := len(grid), len(grid[0])
	// 定义变量 startX、startY、step，分别表示起点的横坐标、纵坐标，以及需要走的总步数
	var startX, startY int
	step := 1

	// 遍历二维数组，找到起点和计算需要走的总步数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
			}

			if grid[i][j] == 0 {
				step++
			}
		}
	}

	// 定义 DFS 函数，实现深度优先搜索和回溯
	var dfs func(int, int, int) int
	dfs = func(i, j, step int) int {
		// 判断当前位置是否越界或者已经访问过，如果是，则返回 0
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == -1 {
			return 0
		}
		// 如果当前位置是终点，表示找到一条符合条件的路径
		if grid[i][j] == 2 {
			if step == 0 { // 步数为 0 时表示找到一条符合条件的路径
				return 1
			}
			return 0
		}

		// 标记当前位置为已访问
		grid[i][j] = -1
		// 定义变量 result，表示从当前位置开始的符合条件的路径数量
		result := 0

		// 分别向上、向下、向左、向右进行深度优先搜索
		result += dfs(i-1, j, step-1)
		result += dfs(i+1, j, step-1)
		result += dfs(i, j-1, step-1)
		result += dfs(i, j+1, step-1)

		// 回溯，将当前位置的标记还原为 0
		grid[i][j] = 0
		return result
	}

	// 调用 DFS 函数，返回符合条件的路径数量
	return dfs(startX, startY, step)
}
