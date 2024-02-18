package string

// reverseString 344. 反转字符串 https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
