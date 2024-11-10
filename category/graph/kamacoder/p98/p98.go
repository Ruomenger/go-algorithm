package main

import (
	"fmt"
)

// main 98. 所有可达路径 https://kamacoder.com/problempage.php?pid=1170
func main() {
	nodeNum, edgeNum := 0, 0
	fmt.Scanf("%d %d\n", &nodeNum, &edgeNum)
	martix := make([][]bool, nodeNum)
	for i := 0; i < nodeNum; i++ {
		martix[i] = make([]bool, nodeNum)
	}

	for i := 0; i < edgeNum; i++ {
		node1, node2 := 0, 0
		fmt.Scanf("%d %d\n", &node1, &node2)
		martix[node1-1][node2-1] = true
	}
	res := make([][]int, 0)
	dfs := func(int, []int) {}
	dfs = func(start int, path []int) {
		if start == nodeNum-1 {
			res = append(res, append(path[:0:0], path...))
			return
		}
		for i := 1; i < nodeNum; i++ {
			path = append(path, i+1)
			if martix[start][i] {
				dfs(i, path)
			}
			path = path[:len(path)-1]
		}

	}
	path := make([]int, 0)
	path = append(path, 1)
	dfs(0, path)
	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		for i := 0; i < len(res); i++ {
			for j := 0; j < len(res[i])-1; j++ {
				fmt.Print(res[i][j], " ")
			}
			fmt.Println(res[i][len(res[i])-1])
		}
	}
}
