package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 双指针 迭代
	// pre is Null, cur is head
	// 遍历， cur的next指向pre，然后pre和cur都向后移动
	//中间需要一个临时变量 保存cur的下一节点

	var pre *ListNode // behind用于记录当前反转链表的头节点，初始化为nil
	cur := head
	for cur != nil { // 循环遍历原链表的每个节点
		tmp := cur.Next // 记录当前节点的下一个节点
		cur.Next = pre  // 反转操作
		pre = cur       // 移动 cur是下一个的pre
		cur = tmp       // 移动到下一个节点
	}

	return pre // 返回反转链表的头节点
}

// 递归
// 终止条件是当前结点或当前结点的下一结点为null
//改变结点的指向 head.next.next = head;就是 head 的下一个节点指向head。

func reverseListV2(head *ListNode) *ListNode {
	// 如果链表为空或者只有一个节点，直接返回原链表头节点
	if head == nil || head.Next == nil {
		return head
	}

	// 递归调用 reverseList 函数，得到反转后的链表头节点 pre
	pre := reverseList(head.Next)

	// 将当前节点的下一个节点的 Next 指针指向当前节点，实现反转
	head.Next.Next = head

	// 将当前节点的 Next 指针置为 nil，避免出现循环
	head.Next = nil

	// 返回反转后的链表头节点 pre
	return pre
}
