package greedy

import (
	"math"
	"slices"
	"sort"
	"strconv"
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

// canCompleteCircuit 134. 加油站 https://leetcode.cn/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
	cur, total, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		cur += (gas[i] - cost[i])
		total += (gas[i] - cost[i])
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

// candy hard 135. 分发糖果 https://leetcode.cn/problems/candy/
func candy(ratings []int) int {
	ans := make([]int, len(ratings))
	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			ans[i] = ans[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			ans[i] = max(ans[i], ans[i+1]+1)
		}
	}

	sum := 0
	for i := 0; i < len(ans); i++ {
		sum += ans[i]
	}
	return sum
}

// lemonadeChange 860. 柠檬水找零 https://leetcode.cn/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	num5, num10 := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			num5++
			continue
		}
		if bills[i] == 10 {
			num5 -= 1
			if num5 < 0 {
				return false
			}
			num10++
		} else if bills[i] == 20 {
			if num10 > 0 {
				num10--
				num5 -= 1
				if num5 < 0 {
					return false
				}
			} else {
				num5 -= 3
				if num5 < 0 {
					return false
				}
			}
		}
	}
	return true
}

// reconstructQueue 406. 根据身高重建队列 https://leetcode.cn/problems/queue-reconstruction-by-height/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}
	return people
}

// findMinArrowShots 452. 用最少数量的箭引爆气球 https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			res++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return res
}

// eraseOverlapIntervals 435. 无重叠区间 https://leetcode.cn/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
			ans++
		}
	}
	return ans
}

// partitionLabels 763. 划分字母区间 https://leetcode.cn/problems/partition-labels/
func partitionLabels(s string) []int {
	ans := make([]int, 0)
	posMap := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		posMap[s[i]] = i
	}
	left := 0
	right := posMap[s[0]]
	for i := left; i <= right; i++ {
		if i >= len(s) {
			ans = append(ans, right-left+1)
			break
		}
		if i > right {
			ans = append(ans, right-left+1)
			left = i
			right = posMap[s[i]]
			continue
		}
		right = max(right, posMap[s[i]])
	}
	return ans
}

// merge 56. 合并区间 https://leetcode.cn/problems/merge-intervals/description/
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			right = max(intervals[i][1], right)
		} else {
			ans = append(ans, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

// monotoneIncreasingDigits 738. 单调递增的数字 https://leetcode.cn/problems/monotone-increasing-digits/description/
func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	str := strconv.Itoa(n)
	ss := []byte(str)
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < len(ss); j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minCameraCover 968. 监控二叉树 https://leetcode.cn/problems/binary-tree-cameras/description/
func minCameraCover(root *TreeNode) int {
	result := 0
	var traversal func(node *TreeNode) int
	// 0：节点无覆盖
	// 1：节点有摄像
	// 2：节点有覆盖
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		left := traversal(node.Left)
		right := traversal(node.Right)
		if left == 2 && right == 2 {
			return 0
		}
		if left == 0 || right == 0 {
			result++
			return 1
		}

		if left == 1 || right == 1 {
			return 2
		}
		return 0
	}
	if traversal(root) == 0 {
		result++
	}
	return result
}
