package array

import "math"

// removeElement simple 27. 移除元素
// https://leetcode.cn/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	old := 0
	newIdx := 0
	for ; newIdx < len(nums); newIdx++ {
		if nums[newIdx] == val {
			continue
		}
		nums[old] = nums[newIdx]
		old += 1
	}
	return old
}

// generateMatrix medium 59. 螺旋矩阵 II
// https://leetcode.cn/problems/spiral-matrix-ii/description/
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, bot := 0, n-1
	left, right := 0, n-1
	count := 1
	target := n * n
	for count <= target {
		for i := left; i <= right; i++ {
			matrix[top][i] = count
			count++
		}
		top++
		for i := top; i <= bot; i++ {
			matrix[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bot][i] = count
			count++
		}
		bot--
		for i := bot; i >= top; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}
	return matrix
}

// minSubArrayLen medium 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	length := math.MaxInt
	left := 0
	right := 0
	total := nums[0]
	for right < len(nums)-1 || total >= target {
		if total < target {
			right++
			total += nums[right]
			continue
		}
		length = min(length, right-left+1)
		total -= nums[left]
		left++
	}
	return length
}

// sortedSquares simple 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
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

// merge simple 88. 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	if n == 0 {
// 		return
// 	}
// 	i := 0
// 	j := 0

// 	for i < len(nums1) {
// 		if i >= m {
// 			nums1[i] = nums2[j]
// 			j++
// 			i++
// 		} else if nums1[i] <= nums2[j] {
// 			i++
// 		} else {
// 			nums1[i], nums2[j] = nums2[j], nums1[i]
// 			for k := j; k < n-1 && nums2[k] > nums2[k+1]; k++ {
// 				nums2[k], nums2[k+1] = nums2[k+1], nums2[k]
// 			}
// 			i++
// 		}
// 	}
// }

// removeDuplicates 26. 删除有序数组中的重复项 simple
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description
func removeDuplicates(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
