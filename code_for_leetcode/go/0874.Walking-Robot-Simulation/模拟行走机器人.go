package leetcode

func robotSim(commands []int, obstacles [][]int) int {
	// 贪心算法 局部最优
	// move为朝着哪个方向走，以坐标轴的顺时针排列
	// d来确定方向，因为move是顺时针排列的，所以+1就是右转，-1就是左转
	// set去重
	// 走之前判定，下一步不会是障碍

	obstacleSet := make(map[[2]int]bool)
	if len(obstacles) != 0 {
		// 将障碍物坐标转化为集合，方便快速查询
		for _, obs := range obstacles {
			if (len(obs)) == 0 {
				continue
			}
			obstacleSet[[2]int{obs[0], obs[1]}] = true
		}
	}

	// 定义机器人的四个移动方向：上、右、下、左
	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 代表当前机器人的方向(在move中的index)，0 表示向上，1 表示向右，2 表示向下，3 表示向左
	direction, result := 0, 0

	// 机器人的当前位置坐标
	i, j := 0, 0

	// 循环遍历 commands
	for _, command := range commands {
		switch command {
		case -2:
			// -2 代表向左转向，更新方向
			direction = (direction - 1 + 4) % 4
		case -1:
			// -1 代表向右转向，更新方向
			direction = (direction + 1) % 4
		default:
			// 其他整数代表机器人向当前方向移动的步数
			x, y := move[direction][0], move[direction][1]
			for k := 0; k < command; k++ {
				// 检查下一步是否有障碍物
				nextPos := [2]int{i + x, j + y}
				if obstacleSet[nextPos] {
					// 遇到障碍物 退出
					break
				}
				i += x
				j += y
				// 更新最大欧式距离的平方
				result = getMax(result, i*i+j*j)
			}
		}
	}

	// 返回最大欧式距离的平方
	return result
}

// getMax 函数：用于比较两个整数，返回较大的那个
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
