package sliding_window

import "math"

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

// minWindow 76. 最小覆盖子串
// level: 困难
// tag: 滑动窗口
// https://leetcode.cn/problems/minimum-window-substring/description/
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	isCovered := func(cntS, cntT [128]int) bool {
		for i := 'A'; i <= 'Z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		return true
	}
	cntT := [128]int{}
	for i := 0; i < len(t); i++ {
		cntT[t[i]]++
	}
	ans := ""
	cntS := [128]int{}
	for left, right := 0, 0; right < len(s); right++ {
		cntS[s[right]]++
		for isCovered(cntS, cntT) {
			if ans == "" || right-left+1 < len(ans) {
				ans = s[left : right+1]
			}
			cntS[s[left]]--
			left++
		}
	}
	return ans
}

func minWindow2(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 扩大窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
