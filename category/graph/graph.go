package graph

import "slices"

// allPathsSourceTarget 797. 所有可能的路径 https://leetcode.cn/problems/all-paths-from-source-to-target/description/
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	dfs := func(int, []int) {}
	dfs = func(start int, path []int) {
		if start == len(graph)-1 {
			res = append(res, slices.Clone(path))
			return
		}
		for i := 0; i < len(graph[start]); i++ {
			path = append(path, graph[start][i])
			dfs(graph[start][i], path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	path = append(path, 0)
	dfs(0, path)
	return res
}

// numIslands 200. 岛屿数量 https://leetcode.cn/problems/number-of-islands/description/
// 广度优先搜索
func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	dirs := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dfs := func(x, y int) {}
	dfs = func(x, y int) {
		for i := 0; i < 4; i++ {
			x1 := x + dirs[i][0]
			y1 := y + dirs[i][1]
			if x1 < 0 || y1 < 0 || x1 >= len(grid) || y1 >= len(grid[0]) {
				continue
			}
			if !visited[x1][y1] && grid[x1][y1] == '1' {
				visited[x1][y1] = true
				dfs(x1, y1)
			}
		}
	}
	result := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if !visited[x][y] && grid[x][y] == '1' {
				result++
				visited[x][y] = true
				dfs(x, y)
			}
		}
	}
	return result
}

// numIslands 200. 岛屿数量 https://leetcode.cn/problems/number-of-islands/description/
// 广度优先搜索
func numIslands2(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	bfs := func(x, y int) {
		queue := make([][2]int, 0)
		visited[x][y] = true
		queue = append(queue, [2]int{x, y})
		for len(queue) != 0 {
			x1 := queue[0][0]
			y1 := queue[0][1]
			queue = queue[1:]
			if x1-1 >= 0 && grid[x1-1][y1] == '1' && !visited[x1-1][y1] {
				visited[x1-1][y1] = true
				queue = append(queue, [2]int{x1 - 1, y1})
			}
			if x1+1 < len(grid) && grid[x1+1][y1] == '1' && !visited[x1+1][y1] {
				visited[x1+1][y1] = true
				queue = append(queue, [2]int{x1 + 1, y1})
			}
			if y1-1 >= 0 && grid[x1][y1-1] == '1' && !visited[x1][y1-1] {
				visited[x1][y1-1] = true
				queue = append(queue, [2]int{x1, y1 - 1})
			}
			if y1+1 < len(grid[0]) && grid[x1][y1+1] == '1' && !visited[x1][y1+1] {
				visited[x1][y1+1] = true
				queue = append(queue, [2]int{x1, y1 + 1})
			}
		}
	}

	result := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if !visited[x][y] && grid[x][y] == '1' {
				visited[x][y] = true
				result++
				bfs(x, y)
			}
		}
	}
	return result
}
