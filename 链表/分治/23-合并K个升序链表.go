package lc

// 23. 合并 K 个升序链表（困难）
//
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你合并所有链表，并返回一个按升序排列的链表。
//
// 分治思路（另可用最小堆）：
//  dfs(l, r) 表示「把 lists[l..r] 合并后的头节点」，
//  出口 l == r 时区间只剩一条链、无需合并，直接返回它本身；
//  递归处把左右两半的结果交给 mergeTwoLists 做节点级归并，
//  合并结果必须 return 上去，否则分治白做。
//  O(log k) 层递归，每层总代价 O(N)，合计 O(N log k) 时间。

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var dfs func(l, r int) *ListNode
	dfs = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}
		mid := l + (r-l)/2
		return mergeTwoLists(dfs(l, mid), dfs(mid+1, r))
	}
	return dfs(0, len(lists)-1)
}

// mergeTwoLists 合并两条升序链表：逐个摘下较小节点接到结果尾部
func mergeTwoLists(ListA, ListB *ListNode) *ListNode {
	dummyNode := &ListNode{}
	pre := dummyNode
	for ListA != nil && ListB != nil {
		if ListA.Val < ListB.Val {
			pre.Next = ListA
			ListA = ListA.Next
		} else {
			pre.Next = ListB
			ListB = ListB.Next
		}
		pre = pre.Next
	}
	if ListA == nil {
		pre.Next = ListB
	} else {
		pre.Next = ListA
	}
	return dummyNode.Next
}
