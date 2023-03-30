/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	return midDiv(0, len(nums)-1, nums, target)
}

func midDiv(head, tail int, nums []int, target int) int {
	mid := (head + tail) / 2
	if nums[mid] == target {
		return mid
	}
	if tail <= head {
		return -1
	}
	if nums[head] < nums[tail] {
		if target > nums[tail] || target < nums[head] {
			return -1
		}
		if target > nums[mid] {
			return midDiv(mid+1, tail, nums, target)
		}
		return midDiv(head, mid-1, nums, target)
	} else {
		return midDiv(head, mid-1, nums, target) * midDiv(mid+1, tail, nums, target) * (-1)
	}
}

// @lc code=end

