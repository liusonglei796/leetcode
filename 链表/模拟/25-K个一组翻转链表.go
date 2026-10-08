package lc

// 25. K 个一组翻转链表
// 思路：先探 k 个节点，数量不足 k 直接原样返回；
// 够 k 个就「探完节点记住第 k+1 个 → 反转前 k 个 → 把剩余部分递归处理 → 接回去」。
// 不够 k 的尾部保持原序，注意不能整体反转。

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	tail := head
	for i := 1; i < k; i++ { // tail 要停在第 k 个节点，再走 k-1 步
		tail = tail.Next
		if tail == nil {
			return head
		}
	}
	rest := tail.Next
	newHead := reverseN(head, k)
	head.Next = reverseKGroup(rest, k)
	return newHead
}

// reverseN 反转以 head 开头的前 n 个节点，返回新的头节点
func reverseN(head *ListNode, n int) *ListNode {
	var prev *ListNode
	cur := head
	for i := 0; i < n && cur != nil; i++ {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
