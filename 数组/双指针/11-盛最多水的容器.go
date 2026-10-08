package lc

// 11. 盛最多水的容器
// 左右双指针相向移动，每次移动较矮的一侧。
func maxArea(height []int) int {
	// TODO: 实现
	maxVal:=0
	left:=0
	right:=len(height)-1
	for left<right{
		hideth:=height[left]
		if height[left]>height[right]{
             hideth=height[right]
		}
		Val:=(right-left)*hideth
		if maxVal<=Val{
			maxVal=Val
		}
		if height[left]<height[right]{
			left++
		}else{
			right--
		}
	}
	return maxVal
}