package lc

// 206. 反转链表
// 思路：逐个把原链表头节点摘下来接到新链表头部。
// prev 始终指向「已反转部分」的头，遍历结束后 prev 就是新头。

// 迭代：三指针边遍历边反转，O(n) 时间 / O(1) 空间
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 必须先记住原后继，再改指针（Go 的赋值语句右边先求值）
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// 递归：先反转 head 之后的全部，再把 head 挂到最前面
// 思路：reverseList(head.Next) 已经让 head.Next 变成了原链表最后一个节点，
// 于是 head.Next.Next = head 就是把 head 接到反转后的链表头。
func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
