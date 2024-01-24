package leetcode

import (
	"github.com/yishengzhishui/doc_for_go/code_for_leetcode/base"
)

// ListNode define
// 定义 ListNode 别名，与 base 包中的 ListNode 类型是相同的
type ListNode = base.ListNode

// 判断链表中是否存在环
func detectCycle(head *ListNode) *ListNode {

	// 快慢指针 (a为非环的长度，b为环的长度)
	// fast和slow相遇时，fast 走的步数是slow步数的 2倍，并且fast 比 slow多走了 n 个环的长度
	// fast = 2slow, fast = slow + nb => slow = nb
	// 所有走到入口节点需要步数：a+nb,
	// 随后 head 开始一步一步，slow从（相遇的位置继续）一步一步
	// slow和head就可重合，即到环的入口节点 head 走a ,slow 再走a就到了

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
			break
		}
	}

	// 如果快指针为 nil 或快指针的下一个节点为 nil，说明链表不存在环，返回 nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 寻找环的入口点
	for slow != head {
		slow = slow.Next
		head = head.Next
	}

	// 返回环的入口点
	return head
}

func detectCycleMap(head *ListNode) *ListNode {
	// 创建一个 map 用于存储节点的地址和索引
	dic := make(map[*ListNode]int)
	index := 0

	// 遍历链表
	for head != nil {
		// 如果当前节点在 map 中已经存在，则说明存在环
		if _, ok := dic[head]; ok {
			return head
		} else {
			// 否则，将当前节点的地址和索引存入 map 中
			dic[head] = index
			index++
			head = head.Next
		}
	}

	// 遍历结束，未发现环，返回 false
	return nil
}
