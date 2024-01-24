package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用一个切片 stack 作为栈，同时使用 result 切片存储遍历结果。
//在遍历过程中，不断弹出栈顶元素，将其值加入结果列表，
//并按右子树、左子树的顺序入栈。
//这样，通过迭代方式实现了二叉树的前序遍历。

//因为在使用栈（stack）进行迭代实现前序遍历时，栈的特性决定了遍历顺序。
//栈是一种后进先出（Last In, First Out，LIFO）的数据结构，
//而前序遍历的访问顺序是先访问根节点，然后依次访问左子树和右子树。
//在使用栈进行迭代时，我们将右子树先入栈，再将左子树入栈。
//这样，出栈的时候就是先弹出左子树，再弹出右子树，从而保证了前序遍历的顺序。
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root} // 初始化栈，根节点入栈
	var result []int           // 初始化结果列表

	for len(stack) > 0 { // 当栈不为空时循环
		cur := stack[len(stack)-1]   // 获取栈顶元素
		stack = stack[:len(stack)-1] // 出栈

		if cur != nil { // 检查当前节点是否为空
			result = append(result, cur.Val)           // 将当前节点值加入结果列表
			stack = append(stack, cur.Right, cur.Left) // 先右后左，右子节点和左子节点依次入栈
		}
	}

	return result // 返回最终结果列表
}

// 递归
// 前序遍历函数
func preorderTraversalV1(root *TreeNode) []int {
	// 如果根节点为空，返回空切片
	if root == nil {
		return []int{}
	}

	// 初始化结果切片，将根节点的值添加到结果中
	result := []int{root.Val}

	// 递归遍历左子树，将结果拼接到当前结果中
	result = append(result, preorderTraversal(root.Left)...)
	// 递归遍历右子树，将结果拼接到当前结果中
	result = append(result, preorderTraversal(root.Right)...)

	// 返回最终的结果切片
	return result
}
