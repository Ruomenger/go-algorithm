package array

import (
	"math"
	"math/rand"
	"slices"
)

// removeElement simple 27. 移除元素
// https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	old := 0
	newIdx := 0
	for ; newIdx < len(nums); newIdx++ {
		if nums[newIdx] == val {
			continue
		}
		nums[old] = nums[newIdx]
		old += 1
	}
	return old
}

// generateMatrix medium 59. 螺旋矩阵 II
// https://leetcode.cn/problems/spiral-matrix-ii/description/
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, bot := 0, n-1
	left, right := 0, n-1
	count := 1
	target := n * n
	for count <= target {
		for i := left; i <= right; i++ {
			matrix[top][i] = count
			count++
		}
		top++
		for i := top; i <= bot; i++ {
			matrix[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bot][i] = count
			count++
		}
		bot--
		for i := bot; i >= top; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}
	return matrix
}

// minSubArrayLen medium 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	length := math.MaxInt
	left := 0
	right := 0
	total := nums[0]
	for right < len(nums)-1 || total >= target {
		if total < target {
			right++
			total += nums[right]
			continue
		}
		length = min(length, right-left+1)
		total -= nums[left]
		left++
	}
	return length
}

// sortedSquares simple 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	answers := make([]int, len(nums))
	// 二分查找nums中最接近0的数
	left := 0
	right := len(nums) - 1
	for left < right-1 {
		mid := left + (right-left)>>1
		if abs(nums[left]) < abs(nums[right]) {
			// 注意判断后要在同号的情况下才能移动right
			if nums[right]*nums[mid] > 0 {
				right = mid
			} else {
				left = mid
			}
		} else {
			if nums[left]*nums[mid] > 0 {
				left = mid
			} else {
				right = mid
			}
		}
	}
	// 从中间的位置分开看
	for i := 0; right <= len(nums)-1 || left >= 0; i++ {
		if right > len(nums)-1 {
			answers[i] = nums[left] * nums[left]
			left--
			continue
		} else if left < 0 {
			answers[i] = nums[right] * nums[right]
			right++
			continue
		}
		if abs(nums[left]) < abs(nums[right]) {
			answers[i] = nums[left] * nums[left]
			left--
		} else {
			answers[i] = nums[right] * nums[right]
			right++
		}
	}
	return answers
}

// merge simple 88. 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	if n == 0 {
// 		return
// 	}
// 	i := 0
// 	j := 0

// 	for i < len(nums1) {
// 		if i >= m {
// 			nums1[i] = nums2[j]
// 			j++
// 			i++
// 		} else if nums1[i] <= nums2[j] {
// 			i++
// 		} else {
// 			nums1[i], nums2[j] = nums2[j], nums1[i]
// 			for k := j; k < n-1 && nums2[k] > nums2[k+1]; k++ {
// 				nums2[k], nums2[k+1] = nums2[k+1], nums2[k]
// 			}
// 			i++
// 		}
// 	}
// }

// removeDuplicates 26. 删除有序数组中的重复项 simple
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description
func removeDuplicates(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}

// removeDuplicates2 80. 删除有序数组中的重复项 II medium
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// majorityElement 169. 多数元素 simple
// https://leetcode.cn/problems/majority-element/description/
func majorityElement(nums []int) int {
	ans := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			count++
		} else if count == 0 {
			count++
			ans = nums[i]
		} else {
			count--
		}
	}
	return ans
}

// rotate 189. 轮转数组 medium
// https://leetcode.cn/problems/rotate-array/description/
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

// hIndex 274. H 指数 medium
// https://leetcode.cn/problems/h-index/description/
func hIndex(citations []int) int {
	ans := len(citations)
	slices.Sort(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] < ans {
			ans--
		} else {
			break
		}
	}
	return ans
}

// RandomizedSet 380. O(1) 时间插入、删除和获取随机元素 medium
// https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
type RandomizedSet struct {
	data  map[int]int
	array []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:  make(map[int]int),
		array: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = len(this.array)
	this.array = append(this.array, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.data[val]
	if !ok {
		return false
	}
	last := len(this.array) - 1
	this.array[idx] = this.array[last]
	this.data[this.array[idx]] = idx
	this.array = this.array[:last]
	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// productExceptSelf 238. 除自身以外数组的乘积 medium
// https://leetcode.cn/problems/product-of-array-except-self/description/
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	temp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		temp *= nums[i+1]
		ans[i] *= temp
	}
	return ans
}

// romanToInt 13. 罗马数字转整数 simple
// https://leetcode.cn/problems/roman-to-integer/description/
func romanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := romans[s[len(s)-1]]
	for i := 1; i < len(s); i++ {
		x, y := romans[s[i-1]], romans[s[i]]
		if x < y {
			ans -= x
		} else {
			ans += x
		}
	}
	return ans
}

// intToRoman 12. 整数转罗马数字 medium
// https://leetcode.cn/problems/integer-to-roman/description/
func intToRoman(num int) string {
	var value2Symbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := []byte{}
	for _, vs := range value2Symbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

// lengthOfLastWord 58. 最后一个单词的长度 simple
// https://leetcode.cn/problems/length-of-last-word/description/
func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ans == 0 {
				continue
			}
			break
		}
		ans++
	}
	return ans
}

// longestCommonPrefix 14. 最长公共前缀 simple
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := math.MaxInt
	for i := 0; i < len(strs); i++ {
		minLen = min(len(strs[i]), minLen)
	}
	if minLen == 0 {
		return ""
	}
	idx := 0
	for ; idx < minLen; idx++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][idx] != strs[0][idx] {
				return strs[0][:idx]
			}
		}
	}
	return strs[0][:idx]
}

// reverseWords 151. 反转字符串中的单词 medium
// https://leetcode.cn/problems/reverse-words-in-a-string/description/
func reverseWords(s string) string {
	reverse := func(strs []byte) {
		for i := 0; i < len(strs)/2; i++ {
			strs[i], strs[len(strs)-i-1] = strs[len(strs)-1-i], strs[i]
		}
	}
	bytes := []byte(s)
	// 首先移除所有的空格，思路是去除所有的空格，然后在单词之间补加一个空格
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		if bytes[fast] != ' ' {
			if slow > 0 {
				bytes[slow] = ' '
				slow++
			}
			for fast < len(s) && bytes[fast] != ' ' {
				bytes[slow] = bytes[fast]
				slow++
				fast++
			}
		}
	}
	bytes = bytes[:slow]
	reverse(bytes)
	left := 0
	for i := 0; i < len(bytes); i++ {
		if i == len(bytes)-1 {
			reverse(bytes[left:])
		} else if bytes[i] == ' ' {
			reverse(bytes[left:i])
			left = i + 1
		}
	}
	return string(bytes)
}
