/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.59%)
 * Likes:    3086
 * Dislikes: 0
 * Total Accepted:    650K
 * Total Submissions: 837.8K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}

// @lc code=end

