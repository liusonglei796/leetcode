package leetcode

// 25. K 个一组翻转链表
// 思路：先探 k 个节点，数量不足 k 直接原样返回；
// 够 k 个就「探完节点记住第 k+1 个 → 反转前 k 个 → 把剩余部分递归处理 → 接回去」。
// 不够 k 的尾部保持原序，注意不能整体反转。

// ListNode: 单链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseKGroup: K 个一组翻转链表的主入口函数
// 来源: 外部调用传入的链表起始头节点 head 以及每组翻转长度 k
// 用于: 执行探查、局部反转与递归子链缝合，返回重构后的整条链表新头节点
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Step 1.1: 探查当前组是否存在完整的 k 个节点
	// 来源: ReverseKGroup 传入的待处理头指针 head 及步长 k
	// 用于: Step 1.2 校验步数并寻找第 k+1 个边界节点
	curr := head
	for i := 0; i < k; i++ {
		// Step 1.2: 校验探查结果，数量不足 k 直接原样返回（递归基，保持尾部原序）
		// 来源: Step 1.1 中当前步进指针 curr 是否提前为 nil
		// 用于: 终止递归并防止不足 k 的尾部被误翻转
		if curr == nil {
			return head
		}
		curr = curr.Next
	}

	// Step 2.1: 探完节点记住第 k+1 个节点作为后续子链递归入口
	// 来源: Step 1.1 走满 k 步后指针停留的位置 curr
	// 用于: Step 3.1 作为剩余链表递归调用的起始入参
	nextGroupHead := curr

	// Step 2.2: 反转当前组的前 k 个节点 [head, nextGroupHead)
	// 来源: 当前组初始头指针 head 以及 Step 2.1 记录的终止边界 nextGroupHead
	// 用于: Step 3.2 产生当前组的新头节点 newHead（原 head 变为新尾节点）
	newHead := reverseSubRange(head, nextGroupHead)

	// Step 3.1: 把剩余部分递归处理
	// 来源: Step 2.1 记住的第 k+1 个节点 nextGroupHead 及步长 k
	// 用于: Step 3.2 获得后续链表递归翻转后的新头节点 remainedHead
	remainedHead := ReverseKGroup(nextGroupHead, k)

	// Step 3.2: 把后续递归处理结果接回当前组翻转后的尾部
	// 来源: 当前组尾节点（即原 head）与 Step 3.1 返回的 remainedHead
	// 用于: 完成前后两段链表的拓扑缝合，返回当前组新头节点 newHead
	head.Next = remainedHead

	return newHead
}

// reverseSubRange: 反转左闭右开区间 [head, stop) 内的节点
// 来源: Step 2.2 传入的当前组头节点 head 与终止边界 stop
// 用于: 将这 k 个节点反转调头，返回反转后的新头节点（即原第 k 个节点）
func reverseSubRange(head, stop *ListNode) *ListNode {
	// Step 2.2.1: 初始化三指针状态
	// 来源: 入参 head
	// 用于: Step 2.2.2 迭代掉转指针方向
	var prev *ListNode = nil
	curr := head

	// Step 2.2.2: 逐节点调转 Next 指向，直到抵达终止边界 stop
	// 来源: Step 2.2.1 的 prev 与 curr
	// 用于: 逐步完成区间内节点的指针倒置
	for curr != stop {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Step 2.2.3: 返回区间反转后的新头节点
	// 来源: Step 2.2.2 循环结束时的 prev
	// 用于: 提供给 Step 2.2 作为当前组的 newHead
	return prev
}
