package leetcode

import "github.com/yishengzhishui/doc_for_go/code_for_leetcode/base"

// ListNode define
type ListNode = base.ListNode



// ReverseKGroup 是反转 k 个节点的方法
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 递归 + 迭代
	// 根据 k 将链表断开
	// 断开后，如果进行反转，count != k 就是不够 k 的数量，不可以反转
	// 随后是递归，递归中包含了迭代的方法（206 问题的反转链表）

	count, cur := 0, head

	// 遍历链表，统计当前链表中节点的数量
	for cur != nil && count != k {
		count++
		cur = cur.Next
	}

	// 如果当前链表中节点数量等于 k，则进行反转
	if count == k {
		// 递归调用，获取下一段链表反转后的头节点
		// 这个cur是当前链表下一段的头节点
		pre := reverseKGroup(cur, k)

		// 反转当前段链表（参考（206 问题的反转链表））
		for i := 0; i < k; i++ {
			tmp := head.Next
			head.Next = pre
			pre = head
			head = tmp
		}

		// head 现在是当前段链表反转后的头节点
		head = pre
	}

	// 返回反转后的头节点
	return head
}
