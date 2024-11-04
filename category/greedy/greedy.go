package greedy

import "slices"

// findContentChildren 455. 分发饼干 https://leetcode.cn/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	i, j := 0, 0
	cnt := 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			cnt++
			i++
		}
		j++
	}
	return cnt
}

// wiggleMaxLength 376. 摆动序列 https://leetcode.cn/problems/wiggle-subsequence/
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	result := 1
	up := true
	down := true
	for i := 0; i < len(nums)-1; i++ {
		if up && nums[i+1] > nums[i] {
			up = false
			down = true
			result++
		} else if down && nums[i+1] < nums[i] {
			down = false
			up = true
			result++
		}
	}
	return result
}
