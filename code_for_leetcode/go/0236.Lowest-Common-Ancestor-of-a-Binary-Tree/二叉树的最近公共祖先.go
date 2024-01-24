package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// 情况分类：
// 1）q,p在左右树 left right ：都存在则返回root，
// 2) p=root，且q在 root 的左或右子树中；
// 3) q=root，且p在 root 的左或右子树中；

// 递归深度优先搜索，不断地在左右子树中查找目标节点 p 和 q，并在找到最近公共祖先时返回。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归终止条件
	// 如果根节点为 nil、p 或 q，直接返回根节点
	if root == nil || root == p || root == q {
		return root
	}

	// 递归查找左子树和右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右子树的结果都不为 nil，说明当前节点是最近公共祖先，直接返回当前节点。

	// 当左右子树的结果都不为 nil 时，说明在当前节点的左右子树中分别找到了目标节点 p 和 q。
	//这意味着当前节点就是最近公共祖先，因为它是最早包含目标节点的节点。
	if left != nil && right != nil {
		return root
	}

	// 如果左子树包含 p 或 q，则返回左子树的结果
	// 否则返回右子树的结果
	if left != nil {
		return left
	} else {
		return right
	}
}
