package leetcode

import (
	"github.com/yishengzhishui/doc_for_go/code_for_leetcode/base"
)

// ListNode define
// 定义 ListNode 别名，与 base 包中的 ListNode 类型是相同的
type ListNode = base.ListNode

// 判断链表中是否存在环
func hasCycle(head *ListNode) bool {
	// 使用快慢指针，初始时两者均指向链表头节点
	fast := head
	slow := head

	// 遍历链表，当快指针到达链表尾部时，表示没有环
	// 如果有环，快慢指针最终会相遇
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针每次移动两步
		slow = slow.Next      // 慢指针每次移动一步

		// 如果快慢指针相遇，说明链表中存在环
		if fast == slow {
			return true
		}
	}

	// 遍历结束，未发现环，返回 false
	return false
}


func hasCycleMap(head *ListNode) bool {
	// 创建一个 map 用于存储节点的地址和索引
	dic := make(map[*ListNode]int)
	index := 0

	// 遍历链表
	for head != nil {
		// 如果当前节点在 map 中已经存在，则说明存在环
		if _, ok := dic[head]; ok {
			return true
		} else {
			// 否则，将当前节点的地址和索引存入 map 中
			dic[head] = index
			index++
			head = head.Next
		}
	}

	// 遍历结束，未发现环，返回 false
	return false
}
