package other

import "strconv"

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
