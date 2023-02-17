/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	var buttons = make(map[int]string, 10)
	buttons[2] = "abc"
	buttons[3] = "def"
	buttons[4] = "ghi"
	buttons[5] = "jkl"
	buttons[6] = "mno"
	buttons[7] = "pqrs"
	buttons[8] = "tuv"
	buttons[9] = "wxyz"
}

// @lc code=end

