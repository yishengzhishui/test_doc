package leetcode

// 遍历grid的每个元素，如果这个元素=='1'，则将其设置为'#'；
// 随后在对其周围的上下左右是个元素进行判定是否为'1', =='1'则设为'#'；

//bfs时不需要设置visited，因为判定周围元素时，如果grid[x][y] != '1',是不会继续的

// numIslands 函数用于计算给定二维字符数组中的岛屿数量
func numIslandsV1(grid [][]byte) int {
	// 如果输入的二维数组为空，直接返回岛屿数量为 0
	if len(grid) == 0 {
		return 0
	}

	// 获取二维数组的行数和列数
	m, n := len(grid), len(grid[0])

	// 初始化岛屿数量为 0
	count := 0

	// 遍历整个二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 调用 bfs 函数进行广度优先搜索，同时将岛屿数量加一
				bfs(grid, i, j)
				count++
			}
		}
	}

	// 返回计算得到的岛屿数量
	return count
}

// bfs 函数用于进行广优先搜索，标记连接的陆地，并将已访问的陆地标记为 '#'
func bfs(grid [][]byte, i, j int) {
	// 获取二维数组的行数和列数
	m, n := len(grid), len(grid[0])

	// 使用切片作为队列，并将当前位置入队
	queue := [][2]int{{i, j}}

	// 遍历队列，知道这个陆地相邻的没有1了，就返回
	for len(queue) > 0 {
		// 出队列，获取当前位置的坐标
		coords := queue[0]
		queue = queue[1:]

		x, y := coords[0], coords[1]

		// 遍历当前位置的上、下、左、右四个方向
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX, newY := x+dir[0], y+dir[1]

			// 判断新位置是否在合法范围内，并且是陆地（'1'）
			if 0 <= newX && newX < m && 0 <= newY && newY < n &&
				grid[newX][newY] == '1' {
				// 将新位置入队，并标记为已访问过（标记为#）
				queue = append(queue, [2]int{newX, newY})
				grid[newX][newY] = '#'
			}
		}
	}
}

// dfs

// numIslands 函数用于计算岛屿数量
func numIslands(grid [][]byte) int {
	// 如果输入的二维数组为空，直接返回岛屿数量为 0
	if len(grid) == 0 {
		return 0
	}

	// 初始化岛屿数量为 0
	count := 0

	// 遍历整个二维数组
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 如果当前位置是陆地（'1'），则进行深度优先搜索（DFS）
			if grid[i][j] == '1' {
				// 调用 dfs 函数进行深度优先搜索，同时将岛屿数量加一
				dfs(grid, i, j)
				count++
			}
		}
	}

	// 返回计算得到的岛屿数量
	return count
}

// dfs 函数用于进行深度优先搜索，将连接的陆地标记为 '#'
func dfs(grid [][]byte, i, j int) {
	// 判断当前位置是否合法，是否是陆地
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	// 将当前位置标记为 '#'
	grid[i][j] = '#'

	// 递归调用 dfs 函数，搜索上、下、左、右四个方向的陆地
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
