package array

// minSubArrayLen 209. 长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	length := 0
	left := 0
	right := 0
	total := nums[0]
	for right < len(nums)-1 || total >= target {
		if total < target {
			right++
			total += nums[right]
			continue
		}
		if length == 0 {
			length = right - left + 1
		} else if length > right-left+1 {
			length = right - left + 1
		}
		total -= nums[left]
		left++
	}
	return length
}
