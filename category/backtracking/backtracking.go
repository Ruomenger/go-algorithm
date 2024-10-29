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

// combinationSum3 216. 组合总和 III https://leetcode.cn/problems/combination-sum-iii/
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, k)
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		if len(path) == k {
			if sum == 0 {
				tmp := make([]int, k)
				copy(tmp, path)
				ans = append(ans, tmp)
			}
			return
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			dfs(i+1, sum-i)
			path = path[:len(path)-1]
		}
	}
	dfs(1, n)
	return ans
}
