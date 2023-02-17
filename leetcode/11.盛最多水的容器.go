/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if max < area {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

