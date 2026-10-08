package lc

// 283. 移动零
// 快慢指针，快指针找非零元素与慢指针交换。
// 283. 移动零 - 快慢指针单次遍历（交换法）
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			if fast != slow {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
	}
}