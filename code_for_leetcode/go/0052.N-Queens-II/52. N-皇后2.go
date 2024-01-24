package leetcode

func totalNQueens(n int) int {
	count := 0
	helper(&count, n, 0, 0, 0, 0)
	return count
}

func helper(count *int, n, row, cols, pie, na int) {
	if row == n {
		*count++
		return
	}

	// 获取可放置皇后的位置
	bits := (^(cols | pie | na)) & ((1 << n) - 1)

	// 遍历可放置皇后的位置
	for bits != 0 {
		p := bits & -bits
		bits = bits & (bits - 1)

		// 递归调用，更新皇后的位置
		helper(count, n, row+1, cols|p, (pie|p)<<1, (na|p)>>1)
	}
}
