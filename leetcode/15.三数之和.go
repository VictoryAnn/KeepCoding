/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := range nums {
		j := i + 1
		k := len(nums) - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for i < j && j < k {
			if j-1 > i && nums[j] == nums[j-1] || nums[j]+nums[k] < -nums[i] {
				j++
				continue
			}
			if nums[j]+nums[k] > -nums[i] {
				k--
				continue
			}
			res = append(res, []int{nums[i], nums[j], nums[k]})
			j++
			k--
		}
	}
	return res
}

// @lc code=end

