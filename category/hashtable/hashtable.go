package hashtable

import (
	"sort"
)

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

// twoSum 1. 两数之和 https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	count := make(map[int]int)
	for i, num := range nums {
		if index, ok := count[target-num]; ok {
			ans[0] = index
			ans[1] = i
			return ans
		}
		count[num] = i
	}
	return ans
}

// fourSumCount 454. 四数相加 II https://leetcode.cn/problems/4sum-ii/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countMap := make(map[int]int)
	ans := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			countMap[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			ans += countMap[-v3-v4]
		}
	}
	return ans
}

// canConstruct 383. 赎金信 https://leetcode.cn/problems/ransom-note/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	count := [26]int{}
	for _, ch := range magazine {
		count[ch-'a']++
	}
	for _, ch := range ransomNote {
		count[ch-'a']--
	}
	for _, v := range count {
		if v < 0 {
			return false
		}
	}
	return true
}

// threeSum 15. 三数之和 https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
