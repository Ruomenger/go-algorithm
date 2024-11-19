package monotonicstack

// dailyTemperatures 739. 每日温度 https://leetcode.cn/problems/daily-temperatures/description/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			j := len(stack) - 1
			for ; j >= 0; j-- {
				if temperatures[stack[j]] >= temperatures[i] {
					break
				}
				res[stack[j]] = i - stack[j]
			}
			stack = stack[:j+1]
			stack = append(stack, i)
		}
	}
	return res
}

// finalPrices 1475. 商品折扣后的最终价格
// level: 简单
// tag: 单调栈
// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
func finalPrices(prices []int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		for len(stack) != 0 && prices[i] <= prices[stack[len(stack)-1]] {
			prices[stack[len(stack)-1]] -= prices[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}

// nextGreaterElement 496. 下一个更大元素 I https://leetcode.cn/problems/next-greater-element-i/description/
func nextGreaterElement(nums1, nums2 []int) []int {
	num1ToIdx := make(map[int]int, len(nums1)) // 预分配空间
	for idx, num1 := range nums1 {
		num1ToIdx[num1] = idx
	}
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for _, num2 := range nums2 {
		for len(stack) > 0 && num2 > stack[len(stack)-1] {
			ans[num1ToIdx[stack[len(stack)-1]]] = num2
			stack = stack[:len(stack)-1]
		}
		if _, ok := num1ToIdx[num2]; ok {
			stack = append(stack, num2)
		}
	}
	return ans
}

// nextGreaterElements 503. 下一个更大元素 II https://leetcode.cn/problems/next-greater-element-ii/description/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return ans
}

// trap 42. 接雨水 https://leetcode.cn/problems/trapping-rain-water/description/
// 双指针解法
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	maxLeft := make([]int, len(height))
	maxLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}
	maxRight := make([]int, len(height))
	maxRight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		cur := min(maxLeft[i], maxRight[i]) - height[i]
		if cur > 0 {
			sum += cur
		}
	}
	return sum
}

// trap2 42. 接雨水 https://leetcode.cn/problems/trapping-rain-water/description/
// 单调栈解法
func trap2(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

// largestRectangleArea 84. 柱状图中最大的矩形 https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
// 双指针
func largestRectangleArea(heights []int) int {
	minLeftIdx := make([]int, len(heights))
	minRightIdx := make([]int, len(heights))
	minLeftIdx[0] = -1
	for i := 1; i < len(heights); i++ {
		t := i - 1
		for t >= 0 && heights[t] >= heights[i] {
			t = minLeftIdx[t]
		}
		minLeftIdx[i] = t
	}
	minRightIdx[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[t] >= heights[i] {
			t = minRightIdx[t]
		}
		minRightIdx[i] = t
	}
	result := 0
	for i := 0; i < len(heights); i++ {
		result = max(result, heights[i]*(minRightIdx[i]-minLeftIdx[i]-1))
	}
	return result
}

// largestRectangleArea 84. 柱状图中最大的矩形 https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
// 单调栈
func largestRectangleArea2(heights []int) int {
	max := 0
	// 使用切片实现栈
	stack := make([]int, 0)
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)
	// 初始化栈，序号从0开始
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		// 结束循环条件为：当即将入栈元素>top元素，也就是形成非单调递增的趋势
		for heights[stack[len(stack)-1]] > heights[i] {
			// mid 是top
			mid := stack[len(stack)-1]
			// 出栈
			stack = stack[0 : len(stack)-1]
			// left是top的下一位元素，i是将要入栈的元素
			left := stack[len(stack)-1]
			// 高度x宽度
			tmp := heights[mid] * (i - left - 1)
			if tmp > max {
				max = tmp
			}
		}
		stack = append(stack, i)
	}
	return max
}
