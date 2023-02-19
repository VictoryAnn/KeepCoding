/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var res = make([]string, 0)

func letterCombinations(digits string) []string {
	var buttons = make(map[byte]string, 10)
	var tmp []byte
	res = []string{}
	if digits == "" {
		return res
	}
	buttons['2'] = "abc"
	buttons['3'] = "def"
	buttons['4'] = "ghi"
	buttons['5'] = "jkl"
	buttons['6'] = "mno"
	buttons['7'] = "pqrs"
	buttons['8'] = "tuv"
	buttons['9'] = "wxyz"
	backtrace(buttons, []byte(digits), tmp)
	return res
}

func backtrace(buttons map[byte]string, digits []byte, tmp []byte) {
	if len(digits) == 0 {
		t := make([]byte, len(tmp))
		copy(t, tmp)
		res = append(res, string(t))
		return
	}
	if v, ok := buttons[digits[0]]; ok {
		for i := 0; i < len(v); i++ {
			tmp = append(tmp, v[i])
			backtrace(buttons, digits[1:], tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	return
}

// @lc code=end

