package leetcode

import (
	"github.com/yishengzhishui/doc_for_go/code_for_leetcode/base"
)

// ListNode define
type ListNode = base.ListNode

//通过递归实现链表的合并。
//首先比较两个链表的头节点的值，选择较小的值作为合并后链表的头节点，
//并将该头节点的 Next 指针指向递归合并后的链表。
//递归终止条件是其中一个链表为空，此时直接返回另一个非空链表。
func mergeTwoListsV1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 递归，
	//优化算法， l1永远设为最小的
	//l1 l2不为Null才能继续执行

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1

	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}

// 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建虚拟头节点(dummy)，用于简化链表操作
	dummy := &ListNode{Val: -1}
	// pre 表示当前合并链表的尾部节点，初始为虚拟头节点
	pre := dummy

	// 循环遍历两个链表，比较节点值并合并
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l1, l2 = l2, l1
		}
		// 如果 l1 的值小于 l2 的值，将 l1 添加到合并链表，并移动 l1 指针
		pre.Next = l1
		l1 = l1.Next
		// 更新合并链表的尾部节点
		pre = pre.Next
	}

	// 处理剩余的链表部分，如果 l1 不为空，将 l1 添加到合并链表；如果 l2 不为空，将 l2 添加到合并链表
	pre.Next = l1
	if l1 == nil {
		pre.Next = l2
	}

	// 返回合并后链表的头节点（虚拟头节点的下一个节点）
	return dummy.Next
}
