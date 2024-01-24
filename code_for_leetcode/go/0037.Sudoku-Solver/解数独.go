package leetcode

func SolveSudoku(board [][]byte) {
	// 定义行、列、块的集合
	row := make([]map[byte]bool, 9) // 用于记录每一行已经使用过的数字
	col := make([]map[byte]bool, 9) // 用于记录每一列已经使用过的数字
	box := make([]map[byte]bool, 9) // 用于记录每一宫已经使用过的数字

	// 初始化集合
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}

	// 初始化空格列表
	empty := make([][2]int, 0)

	// 遍历数独数组，初始化集合和空格列表
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				// 如果当前位置不是空格，表示有数字，将该数字加入行、列、宫的集合
				value := board[i][j]
				row[i][value] = true
				col[j][value] = true
				box[(i/3)*3+j/3][value] = true
			} else {
				// 如果当前位置是空格，记录该空格的位置
				empty = append(empty, [2]int{i, j})
			}
		}
	}

	// 定义回溯函数
	var backtrace func(int) bool
	backtrace = func(index int) bool {
		if index == len(empty) {
			return true // 所有空格都填充完毕，数独解决
		}

		i, j := empty[index][0], empty[index][1]
		bindex := (i/3)*3 + j/3 // 计算所在宫的索引

		for value := byte('1'); value <= byte('9'); value++ {
			// 如果当前数字在行、列、宫中均未使用过，可以尝试填入
			if !row[i][value] && !col[j][value] && !box[bindex][value] {
				row[i][value] = true
				col[j][value] = true
				box[bindex][value] = true
				board[i][j] = value // 将数字填入数独格中
				if backtrace(index + 1) {
					return true // 继续递归下一个空格
				}
				// 回溯：如果填入的数字导致无解，需要将数字移除，尝试其他数字
				row[i][value] = false
				col[j][value] = false
				box[bindex][value] = false
				board[i][j] = '.' // 将格子重新标记为空
			}
		}
		return false
	}

	// 调用回溯函数，从第一个空格开始填充
	backtrace(0)
}
