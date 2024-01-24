package leetcode

// 与岛屿数量 问题类似
// 1. 将边界上的以及和边界相连的'O'变成'//'
// 2. 将所有的'O'变成'X'，'//'变成'O'

// Solve 是解决方案的主函数，用于修改输入矩阵 board
func solve(board [][]byte) {
	// 1. 将边界上的以及和边界相连的'O'变成'#'
	// 2. 将所有的'O'变成'X'，'#'变成'O'
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	// 步骤1：DFS/BFS 将边界上的'O'及其相邻的'O'标记为'#'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				BFS(i, j, &board)
			}
		}
	}

	// 步骤2：遍历整个矩阵，将'O'变成'X'，'#'变成'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// BFS 是广度优先搜索函数，用于将从位置 (i, j) 开始的'O'及其相邻的'O'标记为'#'
// 这里的board是一个指针
func BFS(i, j int, board *[][]byte) {
	m, n := len(*board), len((*board)[0])

	// 用切片模拟队列
	queue := [][]int{{i, j}}
	(*board)[i][j] = '#'

	for len(queue) > 0 {
		// 出队列
		node := queue[0]
		queue = queue[1:]
		i, j := node[0], node[1]

		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && (*board)[x][y] == 'O' {
				// 入队列
				queue = append(queue, []int{x, y})
				(*board)[x][y] = '#'
			}
		}
	}
}

// DFS 是深度优先搜索函数，用于将从位置 (i, j) 开始的'O'及其相邻的'O'标记为'#'
func DFS(i, j int, board *[][]byte) {
	// 检查当前位置是否在矩阵范围内且是'O'
	if i >= 0 && i < len(*board) && j >= 0 && j < len((*board)[0]) && (*board)[i][j] == 'O' {
		// 将当前位置标记为'#'
		(*board)[i][j] = '#'

		// 定义四个方向的偏移量
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		// 对当前位置的四个相邻位置进行DFS
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			DFS(x, y, board)
		}
	}
}
