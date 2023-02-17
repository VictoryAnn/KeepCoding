/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var max = 0
	var mb1 = 0
	var mb2 = 0
	for i := range s {
		length1 := expandFromCenter(s, i, i)
		length2 := expandFromCenter(s, i, i+1)
		if length2 > length1 {
			length1 = length2
		}
		if max < length1 {
			max = length1
			mb1 = i - (length1-1)/2
			mb2 = i + length1/2
		}
	}
	return string(s[mb1+1 : mb2])
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left + 1
}

// @lc code=end

