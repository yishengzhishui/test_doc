package leetcode

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 后序遍历函数
// 递归
func postorder(root *Node) []int {
	// 如果根节点为空，直接返回空切片
	if root == nil {
		return []int{}
	}

	// 初始化结果切片
	result := []int{}

	// 遍历根节点的子节点，对每个子节点进行后序遍历
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}

	// 将当前根节点的值添加到结果切片
	result = append(result, root.Val)

	// 返回最终的结果切片
	return result
}

// 后序遍历 N 叉树的函数
func postorderV1(root *Node) []int {
	stack, result := []*Node{root}, []int{}

	// 循环条件：栈不为空
	for len(stack) > 0 {
		// 弹出栈顶节点
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果当前节点不为空，将其值添加到结果切片
		if cur != nil {
			result = append(result, cur.Val)

			// 将当前节点的子节点入栈
			stack = append(stack, cur.Children...)
		}
	}

	// 将结果切片反转，得到最终的后序遍历序列
	return reverse(result)
}

// reverse 是一个反转切片元素的函数
// 参数 arr 是一个整数切片
// 返回值是反转后的整数切片
func reverse(arr []int) []int {
	// 获取切片的长度
	i, j := 0, len(arr)-1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		j--
		i++
	}

	// 返回反转后的切片
	return arr
}
