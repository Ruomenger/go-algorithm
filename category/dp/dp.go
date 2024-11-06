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
