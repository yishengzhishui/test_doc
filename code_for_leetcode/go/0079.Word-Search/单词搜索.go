package leetcode

// 回溯算法
// 走过的元素可以设为'0'，
// 每次匹配当前元素与word的首字母，递归时传入下层时word将已匹配的首字母剔除
// 递归结束条件 单词找到，或是 路径出界或者矩阵中的值不是word的首字母，返回False

// exist 是入口函数，用于在给定矩阵中查找是否存在与指定单词匹配的路径
func exist(board [][]byte, word string) bool {
	// 如果矩阵为空，直接返回 false
	if len(board) == 0 {
		return false
	}

	// 遍历矩阵中的每个元素，寻找路径的起点
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// 如果在当前位置找到匹配路径，返回 true
			if dfs(board, word, i, j) {
				return true
			}
		}
	}

	// 若遍历完整个矩阵都未找到匹配路径，返回 false
	return false
}

// dfs 是深度优先搜索的主要逻辑
func dfs(board [][]byte, word string, i, j int) bool {
	// 如果单词为空，表示已找到匹配路径，返回 true
	if len(word) == 0 {
		return true
	}

	// 如果当前位置超出矩阵边界，或者当前字符不匹配单词的首字符，返回 false
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[0] != board[i][j] {
		return false
	}

	// 保存当前位置的字符，并将该位置标记为已访问，避免重复访问
	tmp := board[i][j]
	board[i][j] = '0'

	// 递归地探索上、下、左、右四个方向，查找是否存在匹配剩余单词的路径
	result := dfs(board, word[1:], i+1, j) || dfs(board, word[1:], i-1, j) ||
		dfs(board, word[1:], i, j+1) || dfs(board, word[1:], i, j-1)

	// 恢复当前位置的字符，以便在回溯时保持矩阵原始状态
	board[i][j] = tmp

	// 返回是否找到匹配路径
	return result
}
