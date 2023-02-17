/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// 1. 最简单的就是两个for循环
// 2. 由于要返回的是数组下标，所以不能排序，如果排序就要带着坐标。
// 3. 使用map记录所有的数字和坐标，快速查询
func twoSum(nums []int, target int) []int {
	var vals = make(map[int][]int, 0)
	for i, n := range nums {
		if vals[n] == nil {
			vals[n] = []int{i}
		} else {
			vals[n] = append(vals[n], i)
		}
	}
	for i, n := range nums {
		tmp := vals[target-n]
		if len(tmp) == 0 {
			continue
		}
		for _, k := range tmp {
			if k != i {
				return []int{k, i}
			}
		}
	}
	return nil
}

// @lc code=end

