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
