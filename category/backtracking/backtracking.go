package backtracking

// combine 77. 组合 https://leetcode.cn/problems/combinations/
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, k)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
