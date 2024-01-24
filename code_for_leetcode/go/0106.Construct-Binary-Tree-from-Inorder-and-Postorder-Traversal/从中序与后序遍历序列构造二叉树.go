package leetcode

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 如果中序或后序遍历为空，则返回 nil
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 后序遍历的最后一个元素为根节点值
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的索引
	midIndex := indexOf(inorder, rootVal)

	// 递归构建左子树和右子树
	root.Left = buildTree(inorder[:midIndex], postorder[:midIndex])
	root.Right = buildTree(inorder[midIndex+1:], postorder[midIndex:len(postorder)-1])

	return root
}

// 辅助函数，返回元素在切片中的索引
func indexOf(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func buildTreeV2(inorder []int, postorder []int) *TreeNode {
	// 如果后序遍历或中序遍历为空，返回空节点
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 创建中序遍历元素到索引的映射
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}
		fmt.Println(postorder)
		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := inorderMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}
