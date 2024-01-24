package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	// 如果根节点为空，返回深度 0
	if root == nil {
		return 0
	}

	// 如果左子树或右子树为空，返回非空子树的深度加 1
	if root.Left == nil || root.Right == nil {
		return getMax(minDepth(root.Left), minDepth(root.Right)) + 1
	}

	// 如果左右子树都不为空，返回左右子树深度的最小值加 1
	return getMin(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 辅助函数，返回两个整数中的最小值
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数，返回两个整数中的最大值
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 叶子节点是指没有子节点的节点。
// 最小深度是root到达最近叶子节点的深度，
// 也就是root到叶子节点的最小路径

func getMinDepth(node *TreeNode) int {
	//递归出口：node为空
	if node == nil {
		return 0
	}
	//递归式：分三种情况
	//情况1：左子树为空，右边不是空（没有到达叶子节点）
	//则最小深度为右子树的最小深度+1
	if node.Left == nil && node.Right != nil {
		return getMinDepth(node.Right) + 1
	}
	//情况2：右子树为空，左边不是空（也没有到达叶子节点）
	//则最小深度为左子树的最小深度+1
	if node.Right == nil && node.Left != nil {
		return getMinDepth(node.Left) + 1
	}
	//情况3：左右子树都不为空或者都为空， 继续递归
	return getMin(getMinDepth(node.Left), getMinDepth(node.Right)) + 1
}
