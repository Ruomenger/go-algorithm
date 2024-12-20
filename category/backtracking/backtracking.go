package backtracking

import (
	"slices"
	"strings"
)

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

// combinationSum2 40. 组合总和 II https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	nums := make([]int, 0)
	slices.Sort(candidates)
	var backtrace func(int, int)
	backtrace = func(idx, sum int) {
		if idx > len(candidates) || sum < 0 {
			return
		}

		if sum == 0 {
			res = append(res, slices.Clone(nums))
			return
		}
		for i := idx; i < len(candidates); i++ {
			// 同一层的选择中，避免后面的数字跟前面的重复
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			nums = append(nums, candidates[i])
			backtrace(i+1, sum-candidates[i])
			nums = nums[:len(nums)-1]
		}
	}
	backtrace(0, target)
	return res
}

// partition 131. 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/
func partition(str string) [][]string {
	isPalindrome := func(str string) bool {
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-1-i] {
				return false
			}
		}
		return true
	}
	ans := make([][]string, 0)
	strs := make([]string, 0)
	var dfs func(int)
	dfs = func(start int) {
		if start == len(str) {
			ans = append(ans, slices.Clone(strs))
		}
		for i := start; i < len(str); i++ {
			s := str[start : i+1]
			if isPalindrome(s) {
				strs = append(strs, s)
				dfs(i + 1)
				strs = strs[:len(strs)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// restoreIpAddresses 93. 复原 IP 地址 https://leetcode.cn/problems/restore-ip-addresses/
func restoreIpAddresses(s string) []string {
	addresses := make([]string, 0)
	address := make([]string, 0, 4)
	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) && len(address) == 4 {
			addresses = append(addresses, strings.Join(address, "."))
			return
		}
		if len(address) > 4 || (len(address) == 4 && start < len(s)) {
			return
		}

		for i := start; i < len(s); i++ {
			if s[i] == '0' {
				if i == start {
					address = append(address, "0")
					dfs(i + 1)
					address = address[:len(address)-1]
					return
				} else {
					if i-start+1 > 3 {
						continue
					} else if i-start+1 == 3 && s[start] >= '3' {
						continue
					}
					address = append(address, string(s[start:i+1]))
					dfs(i + 1)
					address = address[:len(address)-1]
				}
				continue
			}
			if i-start+1 < 3 || ((i-start+1) == 3 && string(s[start:i+1]) <= "255") {
				address = append(address, string(s[start:i+1]))
				dfs(i + 1)
				address = address[:len(address)-1]
			}
		}
	}
	dfs(0)
	return addresses
}

// subsets 78. 子集 https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(idx int) {
		ans = append(ans, slices.Clone(path))
		if idx == len(nums) {
			return
		}
		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

// subsetsWithDup 90. 子集 II https://leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(start int) {
		ans = append(ans, slices.Clone(path))
		if start == len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

// findSubsequences 491. 非递减子序列 https://leetcode.cn/problems/non-decreasing-subsequences/
func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(int)
	dfs = func(start int) {
		if len(path) >= 2 {
			ans = append(ans, slices.Clone(path))
		}
		if start == len(nums) {
			return
		}
		used := make(map[int]bool, len(nums))
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || path[len(path)-1] <= nums[i] {
				path = append(path, nums[i])
				used[nums[i]] = true
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// permute 46. 全排列 https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make(map[int]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, slices.Clone(path))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}
	dfs()
	return ans
}

// permuteUnique 47. 全排列 II https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make(map[int]bool, len(nums))
	slices.Sort(nums)
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, slices.Clone(path))
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				dfs()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return ans
}

// solveNQueens 51. N 皇后 https://leetcode.cn/problems/n-queens/
func solveNQueens(n int) [][]string {
	var ans [][]string
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	isValid := func(row, col int, chessboard [][]string) bool {
		for i := 0; i < row; i++ {
			if chessboard[i][col] == "Q" {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if chessboard[i][j] == "Q" {
				return false
			}
		}
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if chessboard[i][j] == "Q" {
				return false
			}
		}
		return true
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			ans = append(ans, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(row, i, chessboard) {
				chessboard[row][i] = "Q"
				backtrack(row + 1)
				chessboard[row][i] = "."
			}
		}
	}
	backtrack(0)
	return ans
}

// solveSudoku 37. 解数独 https://leetcode.cn/problems/sudoku-solver/
func solveSudoku(board [][]byte) {
	isValid := func(row, col int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == num {
				return false
			}
			if board[i][col] == num {
				return false
			}
		}
		startRow := (row / 3) * 3
		startCol := (col / 3) * 3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == num {
					return false
				}
			}
		}
		return true
	}

	var backtrace func(int, int) bool
	backtrace = func(x, y int) bool {
		if x == 9 {
			return true
		}
		if board[x][y] != '.' {
			return backtrace(x+(y+1)/9, (y+1)%9)
		}
		for i := byte('1'); i <= '9'; i++ {
			// 先检查再放数就不用排除同位置了
			if isValid(x, y, i) {
				board[x][y] = i
				if backtrace(x+(y+1)/9, (y+1)%9) {
					return true
				}
				board[x][y] = '.'
			}
		}
		return false
	}
	backtrace(0, 0)
}
