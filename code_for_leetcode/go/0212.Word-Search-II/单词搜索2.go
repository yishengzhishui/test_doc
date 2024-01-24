package leetcode

type Trie struct {
	children [26]*Trie // Trie 树节点数组，用于表示 26 个小写字母
	word     string    // 如果当前节点代表一个单词的结尾，则存储该单词
}

// Insert 方法用于将单词插入 Trie 树
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{} // 创建新的 Trie 节点
		}
		node = node.children[ch]
	}
	node.word = word // 将当前节点标记为单词的结尾
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上、下、左、右四个方向的偏移量数组

// findWords 函数实现了在二维字符数组中查找能够组成给定单词列表的单词
func findWords2(board [][]byte, words []string) []string {
	t := &Trie{} // 创建 Trie 树的根节点
	for _, word := range words {
		t.Insert(word) // 将单词插入 Trie 树
	}
	m, n := len(board), len(board[0])
	ans := []string{} // 用于存储结果的字符串数组

	// dfs 函数用于深度优先搜索并查找能够组成单词的路径
	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.children[ch-'a'] // 移动到 Trie 树中的下一个节点
		if node == nil {
			return // 如果 Trie 中没有相应的节点，直接返回
		}
		if node.word != "" {
			ans = append(ans, node.word) // 将找到的单词添加到结果数组中
			node.word = ""               // 清空 word，避免重复添加
		}

		board[x][y] = '#' // 将当前位置标记为已访问
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				dfs(node, nx, ny) // 递归调用 dfs，继续探索下一个位置
			}
		}
		board[x][y] = ch // 恢复当前位置的状态，回溯
	}

	// 遍历二维数组，调用 dfs 函数，查找能够组成单词的路径
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	return ans // 返回结果数组
}

// 采用79的代码也能解决，只是耗时较长
func findWords(board [][]byte, words []string) []string {
	var result []string
	for _, word := range words {
		res := exist(board, word)
		if res {
			result = append(result, word)
		}
	}
	return result
}

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
