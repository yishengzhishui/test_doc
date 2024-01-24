package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalV1(root *TreeNode) []int {
	// 如果根节点为空，返回空切片
	if root == nil {
		return []int{}
	}

	// 初始化结果切片，将根节点的值添加到结果中
	var result []int

	// 递归遍历左子树，将结果拼接到当前结果中
	result = append(result, postorderTraversal(root.Left)...)
	// 递归遍历右子树，将结果拼接到当前结果中
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)

	// 返回最终的结果切片
	return result
}

// 迭代
func postorderTraversal(root *TreeNode) []int {
	// 定义栈
	stack := []*TreeNode{root}
	// 定义结果数组
	var result []int

	// 遍历栈
	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			// 将当前节点值插入结果数组
			result = append(result, node.Val)
			// 将左右子节点依次入栈
			stack = append(stack, node.Left, node.Right)
		}
	}

	// 将结果数组翻转并返回
	reverse(result)
	return result
}

// 辅助函数，翻转数组
func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
