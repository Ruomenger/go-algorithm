package main

import "fmt"

// main 99. 岛屿数量 https://kamacoder.com/problempage.php?pid=1171
func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)
	martix := make([][]bool, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		martix[i] = make([]bool, m)
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := 0
			fmt.Scanf("%d", &x)
			if x == 1 {
				martix[i][j] = true
			}
		}
		// 这个本地测试需要带上，提交不要带
		// fmt.Scanf("\n")
	}
	dir := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dfs := func(x, y int) {}
	dfs = func(x, y int) {
		for i := 0; i < 4; i++ {
			x1 := x + dir[i][0]
			y1 := y + dir[i][1]
			if x1 < 0 || y1 < 0 || x1 >= n || y1 >= m {
				continue
			}
			if !visited[x1][y1] && martix[x1][y1] {
				visited[x1][y1] = true
				dfs(x1, y1)
			}
		}
	}
	result := 0
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if !visited[x][y] && martix[x][y] {
				result++
				visited[x][y] = true
				dfs(x, y)
			}
		}
	}
	fmt.Println(result)
}
