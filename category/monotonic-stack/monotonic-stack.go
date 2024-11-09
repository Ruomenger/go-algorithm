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
