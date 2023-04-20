/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tmp []int

	backtrace(candidates, tmp, 0, target, &res)
	return res
}

func backtrace(candidates, tmp []int, sum, target int, res *[][]int) {
	if sum == target {
		t := make([]int, len(tmp))
		copy(t, tmp)
		(*res) = append((*res), t)
		return
	}
	if sum > target {
		return
	}

	for i, n := range candidates {
		tmp = append(tmp, n)
		sum += n
		backtrace(candidates[i:], tmp, sum, target, res)
		sum -= n
		tmp = tmp[:len(tmp)-1]
	}
}

// @lc code=end

