package string

import (
	"strings"
)

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

// reverseStr 541. 反转字符串 II https://leetcode.cn/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// repeatedSubstringPattern 459. 重复的子字符串 https://leetcode.cn/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	ss := s[1:] + s[:len(s)-1]
	return strings.Contains(ss, s)
}
