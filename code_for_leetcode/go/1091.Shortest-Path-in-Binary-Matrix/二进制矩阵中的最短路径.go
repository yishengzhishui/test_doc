package leetcode

var dx = []int{0, 1, 1, 1, 0, -1, -1, -1, -1}
var dy = []int{1, 1, 0, -1, -1, -1, 0, 1}

func shortestPathBinaryMatrix(grid [][]int) int {
	// BFS：同一层级的 step 是相同的，求的是最少扩散的次数
	m := len(grid)
	// 检查是否为空或者起始点和终点不可达
	if m == 0 || grid[0][0] == 1 || grid[m-1][m-1] == 1 {
		return -1
	} else if m <= 2 {
		return m
	}

	step := 1                           // 步数，初始为1
	queue := make([][2]int, 0)          // 使用数组作为坐标点，作为队列存储
	queue = append(queue, [2]int{0, 0}) // 将起始点加入队列

	for len(queue) > 0 {
		step++             // 进入下一层，步数增加
		size := len(queue) // 当前层的节点数

		for i := 0; i < size; i++ {
			current := queue[0] // 取队首元素
			queue = queue[1:]   // 出队

			for k := 0; k < 8; k++ {
				x, y := current[0]+dx[k], current[1]+dy[k]

				// 如果达到终点，返回步数
				if x == m-1 && y == m-1 {
					return step
				}
				// 判断新坐标是否在合法范围内且为可通行的0
				if 0 <= x && x < m && 0 <= y && y < m && grid[x][y] == 0 {
					queue = append(queue, [2]int{x, y}) // 将新坐标加入队列
					grid[x][y] = -1                     // 标记为已访问，避免重复访问
				}
			}
		}
	}
	return -1 // 如果队列为空仍未找到终点，返回-1表示不可达
}
