package leetcode

import "github.com/yishengzhishui/doc_for_go/code_for_leetcode/base"

type ListNode = base.ListNode

func swapPairs(head *ListNode) *ListNode {
	// 创建虚拟头节点 dummy，并将其 Next 指向原链表头节点
	dummy := &ListNode{Val: 0}
	pre := dummy
	pre.Next = head

	// 循环遍历链表，每次处理两个节点
	for pre.Next != nil && pre.Next.Next != nil {
		first, second := pre.Next, pre.Next.Next

		// 将 pre 的 Next 指向 second，实现交换
		pre.Next = second

		// 保存 second 的下一个节点，防止断链
		tmp := second.Next

		// 将 first 的 Next 指向 tmp，实现交换
		first.Next = tmp

		// 将 second 的 Next 指向 first，实现交换
		second.Next = first

		// 移动 pre 到下一组待交换的节点前的位置
		pre = first
	}

	// 返回虚拟头节点的 Next，即交换后的链表头节点
	return dummy.Next
}

func swapPairsV2(head *ListNode) *ListNode {
	// 终止条件是 head 或 head.Next 是 null
	// 每步只操作 head 和 head.Next，并且 head 指向下面的节点
	if head == nil || head.Next == nil {
		return head // 如果 head 为空或只有一个节点，直接返回 head
	}

	// 取出待交换的两个节点，分别为 first 和 second
	first, second := head, head.Next

	// 递归调用 swapPairsV2 函数，交换后续节点，并将结果赋给 first.Next
	first.Next = swapPairsV2(second.Next)

	// 将 second 的 Next 指向 first，实现节点交换
	second.Next = first

	// 返回 second，即交换后的链表头节点
	return second
}
