package hashtable

// isAnagram 242. 有效的字母异位词 https://leetcode.cn/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for i := range counts {
		if counts[i] != 0 {
			return false
		}
	}
	return true
}

// commonChars 1002. 查找共用字符 https://leetcode.cn/problems/find-common-characters/
func commonChars(words []string) []string {
	myMin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var count [26]int
	for i := range count {
		count[i] = 10000
	}
	ans := make([]string, 0)
	for _, word := range words {
		cur := [26]int{}
		for i := range word {
			cur[word[i]-'a']++
		}
		for i := range cur {
			count[i] = myMin(count[i], cur[i])
		}
	}
	for i := range count {
		for j := 0; j < count[i]; j++ {
			ans = append(ans, string(rune('a'+i)))
		}
	}
	return ans
}
