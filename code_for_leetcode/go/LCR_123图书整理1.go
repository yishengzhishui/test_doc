package leetcode

// 定义 ListNode 结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表函数
func reverseList(head *ListNode) []int {
	result := make([]int, 0)

	// 如果 head 为空链表，直接返回空结果数组
	if head == nil {
		return result
	}

	// 遍历链表，将节点值添加到结果数组中
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	// 使用循环交换数组元素实现数组反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// 返回反转后的结果数组
	return result
}

// 递归
func reverseBookList(head *ListNode) []int {
	// 如果 head 为空链表，直接返回空数组
	if head == nil {
		return []int{}
	}

	// 递归调用 reversePrint 函数，将结果数组与当前节点值拼接
	return append(reverseBookList(head.Next), head.Val)
}
