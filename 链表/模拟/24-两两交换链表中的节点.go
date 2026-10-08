package lc

// 24. 两两交换链表中的节点
// 给你链表的节点链表 head ，两两交换 head 所链接的节点并返回链表。
// 思路：dummy 头 + 「两两一组」推进指针。
// 每轮把 prev.Next 和 prev.Next.Next 交换位置，然后 prev 跳到本组第二个节点。
// 注意：节点是原地交换（重新接指针），不新建节点，空间 O(1)。
//
// 正确写法
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next   // A
		second := first.Next // B
		rest := second.Next  // 后面剩下的部分

		first.Next = rest   // A -> rest
		second.Next = first // B -> A
		prev.Next = second  // prev -> B（B 成了本组新头）

		prev = first // prev 停在本组新尾（原来的 A）
	}
	return dummy.Next
}

// 本次出错的原因（写下来提醒自己）
//
// 出错写法：
//	for prev.Next != nil && prev.Next.Next != nil {
//		temp1 := prev.Next.Next.Next  // rest
//		temp2 := prev.Next            // first
//		prev.Next = prev.Next.Next    // prev.Next = second
//		temp2.Next = temp1            // first.Next = rest
//		prev = (*invalid type)(&temp1) // ← 编译错误
//	}
//
// 1) 编译错误：(*invalid type) 不是合法写法；而且 *ListNode 本身就是指针类型，
//    不能对它取地址（&temp1 得到 **ListNode），要指向 first 直接写 temp2。
// 2) 逻辑漏了一句 —— B.Next = A 没接上：
//    prev.Next = prev.Next.Next 让 B 成了本组新头，
//    但 B.Next 仍然指着原来的后继，A 被摘出链表，
//    结果「每组的第一个节点都丢失」。
// 3) 实测后果：
//		[1,2]        -> 2          （应为 2->1）
//		[1,2,3,4]    -> 2->3       （应为 2->1->4->3）
//		[1,2,3,4,5]  -> 2->3->5    （应为 2->1->4->3->5）
// 4) 记法：交换相邻 A、B 一共要接三次指针：
//		B -> A、A -> rest、prev -> B。
//    交换后 B 是新头、A 是新尾，所以 prev 要停在 A（原来的 first）。
