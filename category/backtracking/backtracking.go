package backtracking

import "slices"

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

// letterCombinations 17. 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	digit2letter := make(map[byte]string)
	digit2letter['2'] = "abc"
	digit2letter['3'] = "def"
	digit2letter['4'] = "ghi"
	digit2letter['5'] = "jkl"
	digit2letter['6'] = "mno"
	digit2letter['7'] = "pqrs"
	digit2letter['8'] = "tuv"
	digit2letter['9'] = "wxyz"
	var dfs func(idx int, str string)
	dfs = func(idx int, str string) {
		if idx == len(digits) {
			if str != "" {
				ans = append(ans, str)
			}
			return
		}
		letters := digit2letter[digits[idx]]
		for _, ch := range letters {
			dfs(idx+1, str+string(ch))
		}
	}
	dfs(0, "")
	return ans
}

// combinationSum 39. 组合总和 https://leetcode.cn/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	nums := make([]int, 0)
	slices.Sort(candidates)
	var backtrace func(int, int)
	backtrace = func(idx, sum int) {
		if idx >= len(candidates) || sum < 0 {
			return
		}

		if sum == 0 {
			res = append(res, slices.Clone(nums))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum < candidates[i] {
				return
			}
			nums = append(nums, candidates[i])
			backtrace(i, sum-candidates[i])
			nums = nums[:len(nums)-1]
		}
	}
	backtrace(0, target)
	return res
}
