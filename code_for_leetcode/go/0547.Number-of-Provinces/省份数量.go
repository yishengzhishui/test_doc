package leetcode

//DFS
func findCircleNumV1(M [][]int) int {
	// DFS
	// visited[i] 表示第 i 个元素是否被访问过。
	// 找到i，并且遍历其所有的朋友（其朋友就是在同一列或行中==1的元素）

	if len(M) == 0 {
		return 0
	}

	size := len(M)
	count, visited := 0, make(map[int]bool) // 使用 map 表示 visited，key 为元素的索引

	for i := 0; i < size; i++ {
		if !visited[i] {
			dfs(M, visited, i)
			count++
		}
	}

	return count
}

// dfs 函数用于深度优先搜索
func dfs(M [][]int, visited map[int]bool, i int) {
	visited[i] = true // 将当前元素标记为已访问

	for j := 0; j < len(M); j++ {
		// 如果 M[i][j] == 1 且 j 未被访问过，则继续递归访问其朋友
		if M[i][j] == 1 && !visited[j] {
			dfs(M, visited, j)
		}
	}
}

// 并查集
func findCircleNum(M [][]int) int {
	// 如果矩阵 M 为空，则朋友圈数量为 0
	if len(M) == 0 {
		return 0
	}

	// 获取矩阵的大小
	size := len(M)

	// 创建并初始化每个元素的父节点为自身的并查集
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
	}

	// 遍历矩阵，对有关系的元素进行合并操作
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if M[i][j] == 1 {
				union(p, i, j)
			}
		}
	}

	// 使用 set 统计集合数量，即朋友圈数量
	set := make(map[int]bool)
	for i := 0; i < size; i++ {
		set[parent(p, i)] = true
	}

	// 返回朋友圈的数量
	return len(set)
}

// union 函数用于合并两个集合
func union(p []int, i, j int) {
	// 查找两个元素所属集合的代表元素
	p1 := parent(p, i)
	p2 := parent(p, j)

	// 将其中一个集合的代表元素的父节点指向另一个集合的代表元素
	p[p1] = p2
}

// parent 函数用于查找元素所属集合的代表元素
func parent(p []int, i int) int {
	// 初始化根节点为元素自身
	root := i

	// 找到集合的代表元素（根节点）
	for p[root] != root {
		root = p[root]
	}

	// 路径压缩，将路径上的所有节点的父节点直接指向根节点
	for p[i] != i {
		x := i
		i = p[i]
		p[x] = root
	}

	// 返回集合的代表元素（根节点）
	return root
}
