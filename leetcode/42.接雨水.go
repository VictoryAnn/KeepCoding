/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var stack []int
	var area int

	for i := range height {
		k := height[i]
		for len(stack) > 0 && k > height[stack[len(stack)-1]] {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				area = area + (i-left-1)*(min(height[i], height[left])-height[bottom])
			}
		}
		stack = append(stack, i)
	}
	return area
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

