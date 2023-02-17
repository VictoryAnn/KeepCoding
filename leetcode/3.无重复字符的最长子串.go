/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	i := 0
	j := 0
	max := 0
	existByte := make(map[byte]bool, 0)
	for j < len(s) {
		if _, exist := existByte[s[j]]; exist {
			for i < j {
				delete(existByte, s[i])
				if s[i] == s[j] {
					i++
					break
				}
				i++
			}
		}

		existByte[s[j]] = true
		if j-i+1 > max {
			max = j - i + 1
		}
		j++
	}
	return max
}

// @lc code=end

