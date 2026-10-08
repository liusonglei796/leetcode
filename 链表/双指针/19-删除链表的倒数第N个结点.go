package lc

// 19. 删除链表的倒数第 N 个结点
//
// 给你链表 head 和整数 n，请你删除链表中倒数第 n 个节点，返回链表的头节点。
// 进阶：能否用一趟（O(L) 时间、O(1) 空间）解决？
//
// 思路（两趟，O(L) 时间 / O(1) 空间）：
// 倒数第 n 个 = 正数第 (L - n + 1) 个，它的节点在正数第 (L - n) 个（要删它的前驱）。
// 第 1 趟数出长度 L，第 2 趟走到第 (L - n) 个节点，用 dummy 兜住「要删的就是头」的情况。
//
// 另一版（一趟快慢指针错位）也能 O(L)/O(1)：
// dummy + fast 先走 n 步，再 for fast.Next != nil 同步前进。
// 注意配对关系：n 步配 fast.Next != nil，或 n+1 步配 fast != nil；
// 混用（n 步 + fast != nil）会错位一格。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	// 1. 第 1 趟：数出链表长度 L
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 2. 第 2 趟：走到第 (L - n) 个节点（倒数第 n 个节点的前驱）
	pre := dummy
	for i := 0; i < length-n; i++ {
		pre = pre.Next
	}

	// 3. 删除并返回新头节点（dummy 保证删头时也有落脚点）
	pre.Next = pre.Next.Next
	return dummy.Next
}
