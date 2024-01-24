package leetcode

type Node struct {
	Val      int
	Children []*Node
}

// levelOrder 是一个层序遍历树的函数
// 参数 root 表示树的根节点
// 返回值是按层级组织的节点值切片
func levelOrderBFS(root *Node) [][]int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return [][]int{}
	}

	// 初始化结果切片和队列
	result := [][]int{}
	queue := []*Node{root}

	// 循环条件：队列不为空
	for len(queue) > 0 {
		// 当前层级的节点数量
		levelSize := len(queue)
		levelValues := []int{}

		// 遍历当前层级的节点
		for i := 0; i < levelSize; i++ {
			// 出队一个节点
			node := queue[0]
			queue = queue[1:]

			// 将节点值添加到当前层级的切片中
			levelValues = append(levelValues, node.Val)

			// 将子节点入队
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}

		// 将当前层级的节点值切片添加到结果切片中
		result = append(result, levelValues)
	}

	// 返回最终的结果切片
	return result
}

// DFS
// 通过递归的方式，按照每一层的顺序，将节点的值存储到对应的层级切片中，
//最终返回一个按层级组织的节点值切片。
func levelOrder(root *Node) [][]int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return [][]int{}
	}

	// 声明一个嵌套函数 helper，用于递归遍历每一层的节点
	var helper func(level int, node *Node)
	var result [][]int

	// 实现 helper 函数
	helper = func(level int, node *Node) {
		if node == nil {
			return
		}
		// 如果当前层级的切片不存在，初始化为空切片
		if len(result) == level {
			result = append(result, []int{})
		}

		// 将当前节点值添加到对应层级的切片中
		result[level] = append(result[level], node.Val)

		// 遍历当前节点的子节点
		for _, child := range node.Children {
			// 递归调用 helper 函数，进入下一层级
			helper(level+1, child)
		}
	}

	// 调用 helper 函数开始层序遍历
	helper(0, root)

	// 返回最终的结果切片
	return result
}
