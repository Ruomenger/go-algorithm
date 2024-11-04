package greedy

import (
	"math"
	"slices"
)

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

// maxSubArray 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
// 可以使用：动态规划、贪心、分治线段树
func maxSubArray(nums []int) int {
	maxSum := math.MinInt
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxSum = max(sum, maxSum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

// maxProfit 122. 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}
	return profit
}

// canJump 55. 跳跃游戏 https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxIdx := nums[0]
	for i := 0; i <= maxIdx; i++ {
		maxIdx = max(i+nums[i], maxIdx)
		if maxIdx >= len(nums)-1 {
			return true
		}
	}
	return false
}

// jump 45. 跳跃游戏 II https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	if (len(nums)) == 1 {
		return 0
	}
	cur, next, ans := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		next = max(i+nums[i], next)
		if i == cur {
			cur = next
			ans++
			if next >= len(nums)-1 {
				return ans
			}
		}
	}
	return ans
}

// largestSumAfterKNegations 1005. K 次取反后最大化的数组和
func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)

	for i := 0; k > 0 && i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	sum := 0
	minNum := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		minNum = min(nums[i], minNum)
	}
	if k%2 != 0 {
		sum -= (2 * minNum)
	}
	return sum
}
