package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 是一个层序遍历树的函数
// 参数 root 表示树的根节点
// 返回值是按层级组织的节点值切片
func largestValuesBFS(root *TreeNode) []int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return []int{}
	}

	// 初始化结果切片和队列
	var result []int
	queue := []*TreeNode{root}

	// 循环条件：队列不为空
	for len(queue) > 0 {
		// 当前层级的节点数量
		levelSize := len(queue)
		levelValues := queue[0].Val

		// 遍历当前层级的节点
		for i := 0; i < levelSize; i++ {
			// 出队一个节点
			node := queue[0]
			queue = queue[1:]

			// 将节点值添加到当前层级的切片中
			levelValues = getMax(levelValues, node.Val)

			// 将子节点入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
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
func largestValues(root *TreeNode) []int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return []int{}
	}

	// 声明一个嵌套函数 helper，用于递归遍历每一层的节点
	var helper func(level int, node *TreeNode)
	var result []int

	// 实现 helper 函数
	helper = func(level int, node *TreeNode) {
		if node == nil {
			return
		}
		// 如果当前层级的切片不存在，初始化为空切片
		if len(result) == level {
			result = append(result, node.Val)
		}

		// 将当前节点值添加到对应层级的切片中
		result[level] = getMax(result[level], node.Val)

		// 遍历当前节点的子节点
		// 递归调用 helper 函数，进入下一层级
		// 将子节点入队
		if node.Right != nil {
			helper(level+1, node.Right)
		}

		if node.Left != nil {
			helper(level+1, node.Left)
		}

	}

	// 调用 helper 函数开始层序遍历
	helper(0, root)

	// 返回最终的结果切片
	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
