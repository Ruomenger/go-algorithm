package string

// reverseString 344. 反转字符串 https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

// strStr 28. 找出字符串中第一个匹配项的下标 https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) || len(needle) == 0 {
		return -1
	}
	next := make([]int, len(needle))
	for i, j := 1, 0; i < len(needle); {
		if needle[i] == needle[j] {
			j++
			next[i] = j
			i++
		} else if j > 0 {
			j = next[j-1]
		} else {
			next[i] = 0
			i++
		}
	}
	for i, j := 0, 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}
