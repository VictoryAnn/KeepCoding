/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	findPoint(0, len(nums)-1, res, nums, target)
	return res
}

func findPoint(start, end int, res []int, nums []int, target int) {
	mid := (start + end) / 2
	if start > end || target > nums[end] || target < nums[start] {
		return
	}
	if nums[mid] == target {
		if mid == 0 || mid > 0 && nums[mid-1] != target {
			res[0] = mid
		}
		if mid == len(nums)-1 || mid < len(nums)-1 && nums[mid+1] != target {
			res[1] = mid
		}
		if res[0] > -1 && res[1] > -1 {
			return
		}
		findPoint(start, mid-1, res, nums, target)
		findPoint(mid+1, end, res, nums, target)
		return
	}
	if nums[mid] > target {
		findPoint(start, mid-1, res, nums, target)
	} else {
		findPoint(mid+1, end, res, nums, target)
	}
	return
}

// @lc code=end

