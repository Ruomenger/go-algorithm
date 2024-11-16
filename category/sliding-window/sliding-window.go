package sliding_window

// maxVowels 1456. 定长子串中元音的最大数目
// level: 中等
// tag: 滑动窗口
// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
func maxVowels(s string, k int) int {
	if k <= 0 || len(s) < k {
		return 0
	}
	isVowel := func(ch byte) int {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			return 1
		}
		return 0
	}
	left, right := 0, 0
	maxNum, curNum := 0, 0

	for right < len(s) {
		if right-left < k {
			curNum += isVowel(s[right])
			right++
			continue
		}
		maxNum = max(maxNum, curNum)
		curNum -= isVowel(s[left])
		left++
	}
	return max(maxNum, curNum)
}
