package array

// sortedSquares 977. 有序数组的平方 https://leetcode.cn/problems/squares-of-a-sorted-array/description/
func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	answers := make([]int, len(nums))
	// 二分查找nums中最接近0的数
	left := 0
	right := len(nums) - 1
	for left < right-1 {
		mid := left + (right-left)>>1
		if abs(nums[left]) < abs(nums[right]) {
			// 注意判断后要在同号的情况下才能移动right
			if nums[right]*nums[mid] > 0 {
				right = mid
			} else {
				left = mid
			}
		} else {
			if nums[left]*nums[mid] > 0 {
				left = mid
			} else {
				right = mid
			}
		}
	}
	// 从中间的位置分开看
	for i := 0; right <= len(nums)-1 || left >= 0; i++ {
		if right > len(nums)-1 {
			answers[i] = nums[left] * nums[left]
			left--
			continue
		} else if left < 0 {
			answers[i] = nums[right] * nums[right]
			right++
			continue
		}
		if abs(nums[left]) < abs(nums[right]) {
			answers[i] = nums[left] * nums[left]
			left--
		} else {
			answers[i] = nums[right] * nums[right]
			right++
		}
	}
	return answers
}
