package lc

// 2. 两数相加
// 链表是低位在前，所以从头部开始就是从个位开始算。
// 逐位相加：sum = l1.Val + l2.Val + carry，个位写进结果，进位存 carry，
// carry 最多为 1。循环条件要带上 carry != 0（最高位可能进位，如 99 + 1）。
// 时间 O(max(m,n))，空间 O(1)（除结果链表）。

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return dummy.Next
}
