package leetcode

// TreeNode 结构体表示二叉树的节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func maxDepthV1(root *TreeNode) int {
	// 如果根节点为空，返回深度 0
	if root == nil {
		return 0
	}

	// 递归计算左子树深度
	left := maxDepth(root.Left)

	// 递归计算右子树深度
	right := maxDepth(root.Right)

	// 返回左右子树深度的最大值加上当前节点深度（+1），表示当前子树的深度
	return getMax(left, right) + 1
}

// getMax 函数返回两个整数的较大值
func getMax(a int, b int) int {
	// 如果 a 大于 b，返回 a；否则返回 b
	if a > b {
		return a
	}
	return b
}

// 迭代，从root开始
func maxDepth(root *TreeNode) int {
	// 如果根节点为空，返回深度 0
	if root == nil {
		return 0
	}

	// 初始化栈和深度最大值
	stack := [][]interface{}{{root, 1}}
	depthMax := 0

	// 循环条件：栈不为空
	for len(stack) > 0 {
		// 弹出栈顶元素，包括当前节点和深度
		curPair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 提取当前节点和深度信息
		cur := curPair[0].(*TreeNode)
		depth := curPair[1].(int)

		// 如果当前节点不为空，更新深度最大值
		if cur != nil {
			depthMax = getMax(depthMax, depth)

			// 将左右子节点和对应深度入栈
			stack = append(stack, []interface{}{cur.Left, depth + 1})
			stack = append(stack, []interface{}{cur.Right, depth + 1})
		}
	}

	// 返回最大深度
	return depthMax
}
