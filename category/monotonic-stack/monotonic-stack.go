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
