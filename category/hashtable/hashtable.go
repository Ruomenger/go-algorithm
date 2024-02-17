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

// intersection 349. 两个数组的交集 https://leetcode.cn/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	countMap := make(map[int]bool)
	for _, num := range nums1 {
		countMap[num] = true
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if countMap[num] == true {
			countMap[num] = false
			result = append(result, num)
		}
	}
	return result
}

// isHappy 202. 快乐数 https://leetcode.cn/problems/happy-number/
func isHappy(n int) bool {
	getSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}
