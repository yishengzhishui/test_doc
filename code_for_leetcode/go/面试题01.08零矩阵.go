package leetcode

func setZeroes(matrix [][]int) {
	// 获取矩阵的行数和列数
	m, n := len(matrix), len(matrix[0])

	// 用于存储零元素的行和列的位置
	var locations [][]int

	// 遍历整个矩阵，找到零元素的位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				locations = append(locations, []int{i, j})
			}
		}
	}

	// 根据零元素的位置，将相应的行和列置零
	for k := 0; k < len(locations); k++ {
		row, column := locations[k][0], locations[k][1]

		// 将列置零
		for i := 0; i < m; i++ {
			matrix[i][column] = 0
		}

		// 将行置零
		for j := 0; j < n; j++ {
			matrix[row][j] = 0
		}
	}
}
