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

// findAnagrams 438. 找到字符串中所有字母异位词
// level: 中等
// tag: 滑动窗口
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(s) < len(p) {
		return ans
	}
	pCnt := [26]int{}
	for i := 0; i < len(p); i++ {
		pCnt[p[i]-'a']++
	}

	sCnt := [26]int{}
	for left, right := 0, 0; right < len(s); right++ {
		sCnt[s[right]-'a']++
		if right-left+1 < len(p) {
			continue
		}
		if sCnt == pCnt {
			ans = append(ans, left)
		}
		sCnt[s[left]-'a']--
		left++
	}

	return ans
}

// findSubstring 30. 串联所有单词的子串
// level: 困难
// tag: 滑动窗口
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	if len(s) < len(words) {
		return ans
	}
	n := len(s)
	m := len(words)
	w := len(words[0])
	total := make(map[string]int, len(words))
	for i := 0; i < m; i++ {
		total[words[i]]++
	}
	for i := 0; i < w; i++ {
		window := make(map[string]int)
		cnt := 0
		for j := i; j+w <= n; j += w {
			if j-i >= w*m {
				word := s[j-m*w : j-m*w+w]
				window[word]--
				if window[word] < total[word] {
					cnt--
				}
			}
			word := s[j : j+w]
			window[word]++
			if window[word] <= total[word] {
				cnt++
			}
			if cnt == m {
				ans = append(ans, j-(m-1)*w)
			}
		}
	}
	return ans
}
