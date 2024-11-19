package dp

import (
	"math"
)

// fib 509. 斐波那契数 https://leetcode.cn/problems/fibonacci-number/description/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre := 0
	cur := 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = pre + cur
		pre = cur
		cur = ans
	}
	return ans
}

// climbStairs 70. 爬楼梯 https://leetcode.cn/problems/climbing-stairs/description/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	num1 := 1
	num2 := 2
	ans := 2
	for i := 3; i <= n; i++ {
		ans = num1 + num2
		num1 = num2
		num2 = ans
	}
	return ans
}

// minCostClimbingStairs 746. 使用最小花费爬楼梯 https://leetcode.cn/problems/min-cost-climbing-stairs/description/
func minCostClimbingStairs(cost []int) int {
	totalCost := 0
	cost1 := 0
	cost2 := 0
	for i := 2; i <= len(cost); i++ {
		totalCost = min(cost2+cost[i-1], cost1+cost[i-2])
		cost1 = cost2
		cost2 = totalCost
	}
	return totalCost
}

// uniquePaths 62. 不同路径 https://leetcode.cn/problems/unique-paths/description/
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		paths[i][0] = 1
	}
	for i := 1; i < n; i++ {
		paths[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]
}

// uniquePathsWithObstacles 63. 不同路径 II https://leetcode.cn/problems/unique-paths-ii/description/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		paths[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		paths[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[m-1][n-1]
}

// integerBreak 343. 整数拆分 https://leetcode.cn/problems/integer-break/description/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}

// numTrees 96. 不同的二叉搜索树 https://leetcode.cn/problems/unique-binary-search-trees/description/
func numTrees(n int) int {
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + (dp[j-1] * dp[i-j])
		}
	}
	return dp[n]
}

// canPartition 416. 分割等和子集 https://leetcode.cn/problems/partition-equal-subset-sum/description/
func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	if totalSum%2 == 1 {
		return false
	}
	totalSum /= 2
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, totalSum+1)
	}

	for i := nums[0]; i <= totalSum; i++ {
		dp[0][i] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= totalSum; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}

	return dp[len(nums)-1][totalSum] == totalSum
}

// lastStoneWeightII 1049. 最后一块石头的重量 II https://leetcode.cn/problems/last-stone-weight-ii/description/
func lastStoneWeightII(stones []int) int {
	if len(stones) <= 1 {
		return stones[0]
	}
	totalSum := 0
	for i := 0; i < len(stones); i++ {
		totalSum += stones[i]
	}
	sum := totalSum / 2
	dp := make([]int, sum+1)
	for i := 0; i < len(stones); i++ {
		for j := sum; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return totalSum - dp[sum] - dp[sum]
}

// findTargetSumWays 494. 目标和 https://leetcode.cn/problems/target-sum/description/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || -target > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	size := (target + sum) >> 1

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, size+1)
	}
	if nums[0] <= size {
		dp[0][nums[0]] = 1
	}
	dp[0][0] = 1
	zero := float64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero++
		}
		dp[i][0] = int(math.Pow(2, zero))
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= size; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][size]
}

// findMaxForm 474. 一和零 https://leetcode.cn/problems/ones-and-zeroes/description/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zero := 0
		one := 0
		for _, ch := range str {
			if ch == '0' {
				zero++
			} else {
				one++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

// change 518. 零钱兑换 II https://leetcode.cn/problems/coin-change-ii/description/
func change(amount int, coins []int) int {
	if amount <= 0 {
		return 1
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// combinationSum4 377. 组合总和 Ⅳ https://leetcode.cn/problems/combination-sum-iv/description/
func combinationSum4(nums []int, target int) int {
	if target <= 0 {
		return 1
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

// coinChange 322. 零钱兑换 https://leetcode.cn/problems/coin-change/description/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// numSquares 279. 完全平方数 https://leetcode.cn/problems/perfect-squares/description/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}
	return dp[n]
}

// wordBreak 139. 单词拆分 https://leetcode.cn/problems/word-break/description/
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	wordMap := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		wordMap[wordDict[i]] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// rob 198. 打家劫舍 https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// rob2 213. 打家劫舍 II https://leetcode.cn/problems/house-robber-ii/description/
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	r1 := rob(nums[0 : len(nums)-1])
	r2 := rob(nums[1:])
	return max(r1, r2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob3 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/description/
func rob3(root *TreeNode) int {
	var robTree func(root *TreeNode) [2]int
	robTree = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{}
		}
		left := robTree(root.Left)
		right := robTree(root.Right)
		return [2]int{root.Val + left[1] + right[1], max(left[0], left[1]) + max(right[0], right[1])}
	}
	results := robTree(root)
	return max(results[0], results[1])
}

// maxProfit 121. 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

// maxProfit2 122. 买卖股票的最佳时机 II https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit2(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

// maxProfit3 123. 买卖股票的最佳时机 III https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}

// maxProfit4 188. 买卖股票的最佳时机 IV https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
func maxProfit4(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i <= 2*k; i += 2 {
		dp[0][i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j <= 2*k; j += 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
		}
		for j := 2; j <= 2*k; j += 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

// maxProfit5 309. 买卖股票的最佳时机含冷冻期 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
func maxProfit5(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	// 状态0：不持有股票，可能是未买入股票或两天前卖出过股票
	// 状态1：持有股票的状态，今天买入了股票或延续之前买入的状态
	// 状态2：卖出股票第一天的状态
	// 状态3：冷冻期
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i], dp[i-1][3]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][3], dp[len(prices)-1][2], dp[len(prices)-1][0])
}

// maxProfit6 714. 买卖股票的最佳时机含手续费 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
func maxProfit6(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}

// lengthOfLIS 300. 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/description/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(dp[i], result)
	}
	return result
}

// findLengthOfLCIS 674. 最长连续递增序列 https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
func findLengthOfLCIS(nums []int) int {
	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curLen++
			maxLen = max(curLen, maxLen)
		} else {
			curLen = 1
		}
	}
	return maxLen
}

// findLength 718. 最长重复子数组 https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
func findLength(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	maxLen := 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen
}

// longestCommonSubsequence 1143. 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/description/
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	maxLen := 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen
}

// maxUncrossedLines 1035. 不相交的线 https://leetcode.cn/problems/uncrossed-lines/description/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	maxLen := 0
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen
}

// maxSubArray 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

// isSubsequence 392. 判断子序列 https://leetcode.cn/problems/is-subsequence/description/
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

// numDistinct 115. 不同的子序列 https://leetcode.cn/problems/distinct-subsequences/description/
func numDistinct(s string, t string) int {
	len1 := len(s)
	len2 := len(t)
	if len(t) > len(s) {
		return 0
	}
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i <= len1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = 0
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len1][len2] % (1e9 + 7)
}

// minDistance 583. 两个字符串的删除操作 https://leetcode.cn/problems/delete-operation-for-two-strings/description/
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+2, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len1][len2]
}

// minDistance2 72. 编辑距离 https://leetcode.cn/problems/edit-distance/description/
func minDistance2(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len1][len2]
}

// countSubstrings 647. 回文子串 https://leetcode.cn/problems/palindromic-substrings/description/
func countSubstrings(s string) int {
	len1 := len(s)
	dp := make([][]bool, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]bool, len1)
	}
	result := 0
	for i := len1 - 1; i >= 0; i-- {
		for j := i; j < len1; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}
		}
	}
	return result
}

func countSubstrings2(s string) int {
	len1 := len(s)
	result := len1
	for i := 0; i < len1; i++ {
		left := i - 1
		right := i + 1
		for left >= 0 && right < len1 && s[left] == s[right] {
			result++
			left--
			right++
		}
	}
	for i := 0; i < len1-1; i++ {
		left := i
		right := i + 1
		for left >= 0 && right < len1 && s[left] == s[right] {
			result++
			left--
			right++
		}
	}
	return result
}

// longestPalindromeSubseq 516. 最长回文子序列 https://leetcode.cn/problems/longest-palindromic-subsequence/description/
func longestPalindromeSubseq(s string) int {
	len1 := len(s)
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len1)
		dp[i][i] = 1
	}
	for i := len1 - 1; i >= 0; i-- {
		for j := i + 1; j < len1; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len1-1]
}

// longestPalindrome 5. 最长回文子串
// level: 中等
// tag: 动态规划
// https://leetcode.cn/problems/longest-palindromic-substring/description/
func longestPalindrome(s string) string {
	lenS := len(s)
	dp := make([][]bool, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenS)
	}
	maxStr := ""
	for i := lenS - 1; i >= 0; i-- {
		for j := i; j < lenS; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > len(maxStr) {
					maxStr = s[i : j+1]
				}
			}
		}
	}
	return maxStr
}

// longestCommonSubStr 最长公共子串
// level: 中等
// tag: 动态规划
func longestCommonSubStr(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	n := len(s) + 1
	m := len(t) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	maxStr := ""
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > len(maxStr) {
					maxStr = s[i-dp[i][j] : i]
				}
			}
		}
	}
	return maxStr
}
