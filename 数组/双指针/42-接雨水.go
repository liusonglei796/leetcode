

// 42. 接雨水
func trapPeak(height []int) int {
    if len(height) <= 2 {
        return 0
    }

    // 1. 找到全局最高柱子
    peakIdx := 0
    for i, h := range height {
        if h > height[peakIdx] {
            peakIdx = i
        }
    }

    water := 0

    // 2. 左半段：向最高峰推进，短板必定在左边
    leftMax := 0
    for i := 0; i < peakIdx; i++ {
        leftMax = max(leftMax, height[i])
        water += leftMax - height[i]
    }

    // 3. 右半段：向最高峰推进，短板必定在右边
    rightMax := 0
    for i := len(height) - 1; i > peakIdx; i-- {
        rightMax = max(rightMax, height[i])
        water += rightMax - height[i]
    }

    return water
}