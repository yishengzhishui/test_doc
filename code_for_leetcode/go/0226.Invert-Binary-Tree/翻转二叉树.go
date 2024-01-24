package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeV1(root *TreeNode) *TreeNode {
	// 如果根节点为空，直接返回 nil
	if root == nil {
		return nil
	}

	// 递归交换左右子树
	//root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	//下面和上面是等价的，好理解一点
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 迭代，使用栈
func invertTree(root *TreeNode) *TreeNode {
	// 如果根节点为空，直接返回 nil
	if root == nil {
		return nil
	}

	// 使用栈进行迭代翻转二叉树
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			// 交换左右子树
			node.Left, node.Right = node.Right, node.Left

			// 将交换后的左右子树入栈
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}

	return root
}
