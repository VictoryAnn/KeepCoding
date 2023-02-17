/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
// p*ss*is
// psssis
//
func isMatch(s string, p string) bool {
	var dp = make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	match := func(i, j int) bool {
		if j == 0 {
			return false
		}
		if p[i-1] == '.' {
			return true
		}
		return p[i-1] == s[j-1]
	}
	for i := 1; i <= len(p); i++ {
		for j := 0; j <= len(s); j++ {
			switch p[i-1] {
			case '*':
				dp[i][j] = dp[i-2][j] // 不重复且清除上一个元素
				if match(i-1, j) {
					dp[i][j] = dp[i][j] || dp[i][j-1] // 重复一次
				}
			default:
				if match(i, j) {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(p)][len(s)]
}

// @lc code=end

