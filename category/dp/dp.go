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
