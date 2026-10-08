package leetcode

import (
	"reflect"
	"testing"
)

// TestReverseKGroup: 测试 K 个一组翻转链表各种边界情况
// 来源: go test 框架驱动
// 用于: 验证算法在完整分组、不足分组、极端值及空链表下的正确性
func TestReverseKGroup(t *testing.T) {
	// Step 1.1: 准备测试用例表
	// 来源: 业务需求与 LeetCode 25 经典用例
	// 用于: Step 1.2 遍历执行测试
	testCases := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "示例1: 5个节点以2个一组翻转",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "示例2: 5个节点以3个一组翻转（末尾2个不足3个保持原序）",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "边界1: k等于1无需翻转",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "边界2: 节点总数不足k直接原样返回",
			input:    []int{1, 2},
			k:        3,
			expected: []int{1, 2},
		},
		{
			name:     "边界3: 单节点链表",
			input:    []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "边界4: 空链表",
			input:    []int{},
			k:        2,
			expected: []int{},
		},
	}

	// Step 1.2: 循环执行用例并校验结果
	// 来源: Step 1.1 的 testCases
	// 用于: 输出测试结论
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Step 1.2.1: 构建输入链表
			// 来源: tc.input 切片
			// 用于: 传入 ReverseKGroup
			head := buildList(tc.input)

			// Step 1.2.2: 执行核心算法
			// 来源: Step 1.2.1 的 head 与 tc.k
			// 用于: 获得输出链表头节点
			resHead := ReverseKGroup(head, tc.k)

			// Step 1.2.3: 提取输出为切片并对比期望
			// 来源: Step 1.2.2 的 resHead
			// 用于: 校验输出与 tc.expected 是否一致
			got := listToSlice(resHead)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("[%s] 失败! 期望: %v, 实际得到: %v", tc.name, tc.expected, got)
			}
		})
	}
}

// buildList: 辅助函数，将切片转为单链表
// 来源: TestReverseKGroup 中传入的切片 nums
// 用于: 构造单链表头节点
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// listToSlice: 辅助函数，将单链表转为切片方便校验
// 来源: TestReverseKGroup 中传入的链表头节点 head
// 用于: 生成切片便于断言比较
func listToSlice(head *ListNode) []int {
	res := make([]int, 0)
	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
}
