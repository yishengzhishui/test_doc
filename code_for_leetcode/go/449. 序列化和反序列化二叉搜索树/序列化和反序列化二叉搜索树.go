package leetcode

import (
	"strconv"
	"strings"
)

// TreeNode 是二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec 结构用于序列化和反序列化二叉树
type Codec struct{}

// Constructor 返回一个 Codec 实例
func Constructor() Codec {
	return Codec{}
}

//
//// DFS
//// Serialize 将二叉树序列化为字符串
//func (codec *Codec) serialize(root *TreeNode) string {
//	if root == nil {
//		return "X,"
//	}
//
//	// 递归序列化左右子树
//	left := codec.serialize(root.Left)
//	right := codec.serialize(root.Right)
//
//	// 使用逗号分隔每个节点的值，并用 X 表示空节点
//	return strconv.Itoa(root.Val) + "," + left + right
//}
//
//// Deserialize 将字符串反序列化为二叉树
//func (codec *Codec) deserialize(data string) *TreeNode {
//	// 将字符串按逗号分割成字符串数组
//	values := strings.Split(data, ",")
//	return codec.deserializeHelper(&values)
//}
//
//// deserializeHelper 是反序列化的辅助函数
//func (codec *Codec) deserializeHelper(values *[]string) *TreeNode {
//	// 从字符串数组中取出第一个值
//	value := (*values)[0]
//	// 移除已处理的值
//	*values = (*values)[1:]
//
//	// 如果当前值是 X，表示空节点
//	if value == "X" {
//		return nil
//	}
//
//	// 将字符串转换为整数作为当前节点的值
//	val, _ := strconv.Atoi(value)
//	// 创建当前节点
//	root := &TreeNode{Val: val}
//
//	// 递归处理左右子树
//	root.Left = codec.deserializeHelper(values)
//	root.Right = codec.deserializeHelper(values)
//
//	return root
//}

//BFS

// serialize 将二叉树序列化为字符串
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	// 使用队列辅助进行层级遍历
	queue := []*TreeNode{root}
	var result strings.Builder

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			// 将节点值转为字符串并追加到结果中
			result.WriteString(strconv.Itoa(node.Val))
			result.WriteString(",")

			// 将左右子节点加入队列
			queue = append(queue, node.Left, node.Right)
		} else {
			// 空节点用 "X" 表示，并追加到结果中
			result.WriteString("X,")
		}
	}

	return result.String()
}

// deserialize 将字符串反序列化为二叉树
func (codec *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// 拆分字符串为节点值数组
	vals := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(vals[0])

	// 创建根节点
	root := &TreeNode{Val: rootVal}

	// 使用队列辅助进行反序列化
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 处理左子节点
		leftVal := vals[index]
		index++

		if leftVal != "X" {
			leftValInt, _ := strconv.Atoi(leftVal)
			node.Left = &TreeNode{Val: leftValInt}

			// 将左子节点加入队列
			queue = append(queue, node.Left)
		}

		// 处理右子节点
		rightVal := vals[index]
		index++

		if rightVal != "X" {
			rightValInt, _ := strconv.Atoi(rightVal)
			node.Right = &TreeNode{Val: rightValInt}

			// 将右子节点加入队列
			queue = append(queue, node.Right)
		}
	}

	return root
}
