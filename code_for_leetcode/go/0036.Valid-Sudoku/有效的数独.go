package leetcode

func isValidSudoku(board [][]byte) bool {
	// 初始化行、列、九宫格的集合数组
	row := make([]map[byte]bool, 9) // 存储每一行已出现的数字集合
	col := make([]map[byte]bool, 9) // 存储每一列已出现的数字集合
	box := make([]map[byte]bool, 9) // 存储每个九宫格已出现的数字集合

	for i := range row {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}

	// 遍历数独表格
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				value := board[i][j]

				// 检查行、列、九宫格是否有重复数字
				if row[i][value] || col[j][value] || box[(i/3)*3+j/3][value] {
					return false
				}

				// 更新行、列、九宫格的集合
				row[i][value] = true
				col[j][value] = true
				box[(i/3)*3+j/3][value] = true
			}
		}
	}

	return true // 数独表格有效，无重复数字
}
