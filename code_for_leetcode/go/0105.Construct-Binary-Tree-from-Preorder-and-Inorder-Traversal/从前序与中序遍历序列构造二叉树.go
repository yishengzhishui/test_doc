package leetcode

// TreeNode 结构定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 标准递归方法
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 如果前序遍历或中序遍历为空，返回空节点
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 创建根节点，取前序遍历的第一个元素
	root := &TreeNode{Val: preorder[0]}

	// 在中序遍历中找到根节点的索引
	midIndex := indexOf(inorder, root.Val)

	// 递归构建左子树，注意切片的范围
	// 前序（根-左-右）
	// 中序（左-根-右）
	// 所以实际的left在 preorder[1,midIndex+1), inorder的[0,midIndex)
	root.Left = buildTree(preorder[1:midIndex+1], inorder[:midIndex])

	// 递归构建右子树，注意切片的范围
	root.Right = buildTree(preorder[midIndex+1:], inorder[midIndex+1:])

	return root
}

// indexOf 返回目标值在切片中的索引
func indexOf(slice []int, target int) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1
}

// buildTree 根据前序和中序遍历构建二叉树
// 存储inorder的序列 Hash
// 方法 helper 的参数是中序序列中当前子树的左右边界
func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	// 如果前序遍历或中序遍历为空，返回空节点
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 创建中序遍历元素到索引的映射
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	// 调用递归辅助函数构建二叉树
	return buildTreeHelper(preorder, inorderMap, 0, len(inorder)-1)
}

// buildTreeHelper 递归辅助函数，根据前序和中序遍历构建二叉树
func buildTreeHelper(preorder []int, inorderMap map[int]int, left, right int) *TreeNode {
	// 如果左边界大于右边界，返回空节点
	if left > right {
		return nil
	}
	// 如果前序遍历为空，返回空节点
	if len(preorder) == 0 {
		return nil
	}
	// 从前序遍历中取出根节点的值
	rootVal := preorder[0]
	//注意这里要换一下，原来的递归是不会改变preorder的，但是这里会改变
	// 为了不影响底层数据，需要copy一份
	copy(preorder, preorder[1:])
	// 前一步中已经复制了 preorder[1:] 的元素，所以最后一个元素是重复的，需要删除。
	preorder = preorder[:len(preorder)-1]

	// 创建根节点
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的索引
	midIndex := inorderMap[rootVal]

	// 递归构建左子树，注意更新左右边界
	// 递归构建右子树，注意更新左右边界
	root.Right = buildTreeHelper(preorder, inorderMap, midIndex+1, right)

	root.Left = buildTreeHelper(preorder, inorderMap, left, midIndex-1)


	return root
}