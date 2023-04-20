/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var tmp []int
	var res [][]int
	permuteAll(0, nums, tmp, &res)
	return res
}

func permuteAll(index int, nums, tmp []int, res *[][]int) {
	if len(nums) == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		(*res) = append(*res, t)
		return
	}
	for i := range nums {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		tmp = append(tmp, nums[i])
		numsCopy = append(numsCopy[:i], numsCopy[i+1:]...)
		permuteAll(0, numsCopy, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}

// @lc code=end
 
