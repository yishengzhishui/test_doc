package leetcode

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 前序遍历函数
func preorder(root *Node) []int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return []int{}
	}

	// 将当前根节点的值添加到结果切片
	result := []int{root.Val}

	// 遍历根节点的子节点，对每个子节点进行后序遍历
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}

	// 返回最终的结果切片
	return result
}
func preorderV1(root *Node) []int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return []int{}
	}

	// 初始化栈和结果切片
	stack := []*Node{root}
	result := []int{}

	// 循环条件：栈不为空
	for len(stack) > 0 {
		// 弹出栈顶节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果当前节点不为空，将其值添加到结果切片
		if cur != nil {
			result = append(result, cur.Val)

			// 将当前节点的子节点逆序入栈
			for i := len(cur.Children) - 1; i >= 0; i-- {
				stack = append(stack, cur.Children[i])
			}
		}
	}

	// 返回最终的结果切片
	return result
}
