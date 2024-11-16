package other

import (
	"strconv"
	"strings"
)

// maximum69Number 1323. 6 和 9 组成的最大数字
// level: simple
// https://leetcode.cn/problems/maximum-69-number/solutions/
func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	ans := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '6' {
			ans = append(ans, '9')
			ans = append(ans, []byte(str)[i+1:]...)
			break
		} else {
			ans = append(ans, '9')
		}
	}
	maxNum, _ := strconv.Atoi(string(ans))
	return maxNum
}

// printVertically 1324. 竖直打印单词
// level: medium
// https://leetcode.cn/problems/print-words-vertically/description/
func printVertically(s string) []string {
	strs := strings.Split(s, " ")
	ans := make([]string, 0)
	maxLen := 0
	for i := 0; i < len(strs); i++ {
		maxLen = max(maxLen, len(strs[i]))
	}
	for i := 0; i < maxLen; i++ {
		bytes := make([]byte, len(strs))
		for j := 0; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				bytes[j] = ' '
			} else {
				bytes[j] = strs[j][i]
			}
		}
		for j := len(bytes) - 1; j >= 0; j-- {
			if bytes[j] == ' ' {
				bytes = bytes[:j]
			} else {
				break
			}
		}
		ans = append(ans, string(bytes))
	}
	return ans
}
