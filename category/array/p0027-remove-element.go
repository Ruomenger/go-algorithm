package array

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
