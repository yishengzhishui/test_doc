package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	// 返回最终的结果切片
	return result
}

// 中序遍历函数(左根右)
// 将左子树的所有节点依次入栈，然后出栈并处理，再转向右子树，循环进行，最终得到整个树的中序遍历结果。
// 中序遍历函数
func inorderTraversalV2(root *TreeNode) []int {
	// 初始化栈和结果切片
	stack := []*TreeNode{} // 用于存储待访问的节点的栈
	result := []int{}      // 存储最终的中序遍历结果
	cur := root            // 当前节点初始化为根节点

	// 循环条件：栈不为空或当前节点不为空
	for len(stack) > 0 || cur != nil {
		// 将当前节点及其左子树入栈，直到左子树为空
		for cur != nil {
			stack = append(stack, cur) // 当前节点入栈
			cur = cur.Left             // 移动到左子树
		}

		// 弹出栈顶元素，将其值添加到结果切片
		cur = stack[len(stack)-1]        // 弹出栈顶元素
		stack = stack[:len(stack)-1]     // 栈顶元素出栈
		result = append(result, cur.Val) // 将节点值添加到结果切片

		// 将当前节点指向右子树
		cur = cur.Right // 移动到右子树
	}

	// 返回最终的结果切片
	return result
}
