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
