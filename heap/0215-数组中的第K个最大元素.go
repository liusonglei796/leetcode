package heap

import "container/heap"

// 215. 数组中的第K个最大元素
// 中等

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 示例 1:
// 输入: nums = [3,2,1,5,6,4], k = 2
// 输出: 5

// 示例 2:
// 输入: nums = [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4

// 提示:
// 1 <= k <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4

// 最小堆
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
