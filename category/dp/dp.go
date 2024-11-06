package dp

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
