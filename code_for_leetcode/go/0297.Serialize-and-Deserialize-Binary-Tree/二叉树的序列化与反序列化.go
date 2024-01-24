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

// DFS
// Serialize 将二叉树序列化为字符串
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "X,"
	}

	// 递归序列化左右子树
	left := codec.serialize(root.Left)
	right := codec.serialize(root.Right)

	// 使用逗号分隔每个节点的值，并用 X 表示空节点
	return strconv.Itoa(root.Val) + "," + left + right
}

// Deserialize 将字符串反序列化为二叉树
func (codec *Codec) deserialize(data string) *TreeNode {
	// 将字符串按逗号分割成字符串数组
	values := strings.Split(data, ",")
	return codec.deserializeHelper(&values)
}

// deserializeHelper 是反序列化的辅助函数
func (codec *Codec) deserializeHelper(values *[]string) *TreeNode {
	// 从字符串数组中取出第一个值
	value := (*values)[0]
	// 移除已处理的值
	*values = (*values)[1:]

	// 如果当前值是 X，表示空节点
	if value == "X" {
		return nil
	}

	// 将字符串转换为整数作为当前节点的值
	val, _ := strconv.Atoi(value)
	// 创建当前节点
	root := &TreeNode{Val: val}

	// 递归处理左右子树
	root.Left = codec.deserializeHelper(values)
	root.Right = codec.deserializeHelper(values)

	return root
}

