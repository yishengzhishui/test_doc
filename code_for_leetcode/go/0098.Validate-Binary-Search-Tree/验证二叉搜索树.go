package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//二叉搜索树定义如下：
//
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

// 中序遍历，左-根-右就是升序

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isValidBSTV1(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(node *TreeNode, minVal, maxVal int) bool {
	// 如果节点为空，返回 true
	if node == nil {
		return true
	}

	data := node.Val

	// 检查节点值是否在有效范围内
	// 如果中间节点大于当前最大值或小于最小值，则False

	if data >= maxVal || data <= minVal {
		return false
	}

	// 递归检查左子树
	if !helper(node.Left, minVal, data) {
		return false
	}

	// 左子树没有问题，检查右子数
	return helper(node.Right, data, maxVal)
}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{} // 创建一个栈用于迭代
	pre := math.MinInt     // 初始化前一个节点的值为负无穷大
	cur := root            // 从根节点开始遍历

	// 中序遍历，如果不是升序就返回false
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur) // 将当前节点及其左子树入栈
			cur = cur.Left             // 移动到左子树
		}

		cur = stack[len(stack)-1] // 弹出栈顶节点
		stack = stack[:len(stack)-1]

		if pre >= cur.Val {
			return false // 如果中序遍历得到的节点的值小于等于前一个节点的值，说明不是二叉搜索树
		}

		pre = cur.Val   // 更新前一个节点的值
		cur = cur.Right // 移动到右子树
	}

	return true
}
