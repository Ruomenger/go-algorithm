package array

// generateMatrix 59. 螺旋矩阵 II  https://leetcode.cn/problems/spiral-matrix-ii/description/
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
