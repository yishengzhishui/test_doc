package leetcode

// DFS
// 点击一次之后；从点击的位置开始
// 如果点击位置是'M'，即有地雷，则设为'X'后直接返回数组；
// 如果不是'M'，则开始遍历其八方向的元素，进行判定； 如果八方向上有'M'，则count+1；
// 将count结果显示在元素位置后，进行下一步递归计算

func updateBoardV1(board [][]byte, click []int) [][]byte {
	// 获取点击位置的行和列
	i, j := click[0], click[1]

	// 如果点击位置是地雷（'M'），则将其设为 'X' 后直接返回数组
	if board[i][j] == 'M' {
		board[i][j] = 'X'
		return board
	}

	// 定义八个方向的偏移量，用于遍历周围的元素
	direction := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

	// 定义深度优先搜索函数
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 判断当前位置是否合法，如果不是 'E' 则直接返回
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'E' {
			return
		}

		count := 0

		// 遍历八个方向，统计周围地雷数量
		for _, d := range direction {
			dx, dy := i+d[0], j+d[1]
			// 位置合法且这个位置是'M'（地雷）
			if 0 <= dx && dx < len(board) &&
				0 <= dy && dy < len(board[0]) &&
				board[dx][dy] == 'M' {
				count++
			}
		}

		// 如果周围没有地雷，则将当前位置设为 'B'，并递归处理周围的元素
		if count == 0 {
			board[i][j] = 'B'
		} else {
			// 如果周围有地雷，则将当前位置设为地雷数量，并结束递归
			// 将 count 转换为字节并加上 '0'，实际上是将数字转换为对应的字符。
			board[i][j] = byte(count) + '0'
			return
		}

		// 如果这个周围（八个方向）没有地雷
		// 对周围的元素递归调用深度优先搜索
		for _, d := range direction {
			dx, dy := i+d[0], j+d[1]
			if 0 <= dx && dx < len(board) && 0 <= dy && dy < len(board[0]) {
				dfs(dx, dy)
			}
		}
	}

	// 调用深度优先搜索函数
	dfs(i, j)

	// 返回更新后的数组
	return board
}

// BFS
// 点击一次之后；从点击的位置开始
// 如果点击位置是'M'，即有地雷，则设为'X'后直接返回数组；
// 如果不是'M'，则开始遍历其八方向的元素，进行判定； 如果八方向上有'M'，则count+1；
// count == 0 则将其周围八个元素加入到队列中；count != 0将count结果显示在元素位置后
// 为了防止超时，使用visited保存访问过的元素

func updateBoard(board [][]byte, click []int) [][]byte {
	// 检查游戏板是否为空
	if len(board) == 0 {
		return nil
	}

	i, j := click[0], click[1]

	// 如果点击的位置是地雷，直接将其标记为 'X' 并返回
	if board[i][j] == 'M' {
		board[i][j] = 'X'
		return board
	}

	m, n := len(board), len(board[0])
	direction := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	visited := make(map[[2]int]bool)

	// 定义 BFS 函数，用于从点击的位置开始遍历
	var bfs func(a, b int)
	bfs = func(a, b int) {
		// 初始化队列，将点击位置入队
		queue := [][2]int{{a, b}}

		for len(queue) > 0 {
			coords := queue[0]
			queue = queue[1:]
			i, j := coords[0], coords[1]
			count := 0

			// 计算周围地雷的数量
			for _, dir := range direction {
				dx, dy := dir[0]+i, dir[1]+j
				if 0 <= dx && dx < m && 0 <= dy && dy < n && board[dx][dy] == 'M' {
					count++
				}
			}

			// 根据地雷数量进行不同的处理
			if count == 0 {
				// 如果周围没有地雷，将当前位置标记为 'B'，并将周围未访问的位置入队
				board[i][j] = 'B'
				for _, dir := range direction {
					dx, dy := dir[0]+i, dir[1]+j
					if 0 <= dx && dx < m && 0 <= dy && dy < n && !visited[[2]int{dx, dy}] {
						queue = append(queue, [2]int{dx, dy})
						visited[[2]int{dx, dy}] = true
					}
				}
			} else {
				// 如果周围有地雷，将当前位置标记为地雷数量
				board[i][j] = byte(count) + '0'
			}
		}
	}

	// 调用 BFS 函数，从点击的位置开始遍历
	bfs(i, j)

	return board
}
