package heap

import "container/heap"

// 347. 前 K 个高频元素
// 中等

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。
// 你可以按 任意顺序 答案。
// 进阶：你能设计一个时间复杂度为 O(n log n) 的算法解决此问题吗？

// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]

// 示例 2:
// 输入: nums = [1], k = 1
// 输出: [1]

// 提示:
// 1 <= nums.length <= 10^5
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

// 元素频率结构体
type Element struct {
	num  int
	freq int
}

// 最小堆，按频率排序
type freqHeap []Element

func (h freqHeap) Len() int           { return len(h) }
func (h freqHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h freqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}
func (h *freqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 统计频率
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// 用最小堆维护前k个高频元素
	h := &freqHeap{}
	heap.Init(h)

	for num, freq := range count {
		heap.Push(h, Element{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 提取结果
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (*h)[i].num
	}
	return result
}
