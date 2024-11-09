package dp

import (
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2}, 1},
		{"test2", args{3}, 2},
		{"test3", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2}, 2},
		{"test2", args{3}, 3},
		{"test3", args{4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{10, 15, 20}}, 15},
		{"test2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{3, 7}, 28},
		{"test2", args{3, 2}, 3},
		{"test3", args{7, 3}, 28},
		{"test4", args{3, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"test2", args{[][]int{{0, 1}, {0, 0}}}, 1},
		{"test3", args{[][]int{{0, 0}, {0, 1}}}, 0},
		{"test4", args{[][]int{{0, 0}, {1, 1}, {0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2}, 1},
		{"test2", args{10}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.args.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{3}, 5},
		{"test2", args{1}, 1},
		{"test3", args{4}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]int{1, 5, 11, 5}}, true},
		{"test2", args{[]int{1, 2, 3, 5}}, false},
		{"test3", args{[]int{2, 2, 3, 5}}, false},
		{"test3", args{[]int{2, 2, 1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastStoneWeightII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 7, 4, 1, 8, 1}}, 1},
		{"test2", args{[]int{31, 26, 33, 21, 40}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightII(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeightII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
		{"test2", args{[]int{1}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]string{"10", "0001", "111001", "1", "0"}, 5, 3}, 4},
		{"test2", args{[]string{"10", "0", "1"}, 1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, []int{5, 2, 1}}, 4},
		{"test2", args{3, []int{2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3}, 4}, 7},
		{"test2", args{[]int{9}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 5}, 11}, 3},
		{"test2", args{[]int{2}, 3}, -1},
		{"test3", args{[]int{1}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{12}, 3},
		{"test2", args{13}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"leetcode", []string{"leet", "code"}}, true},
		{"test2", args{"applepenapple", []string{"apple", "pen"}}, true},
		{"test3", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 1}}, 4},
		{"test2", args{[]int{2, 7, 9, 3, 1}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 3, 2}}, 3},
		{"test2", args{[]int{1, 2, 3, 1}}, 4},
		{"test3", args{[]int{1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 2}
	root1.Left.Right = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 3}
	root1.Right.Right = &TreeNode{Val: 1}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 4}
	root2.Left.Left = &TreeNode{Val: 1}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 5}
	root2.Right.Right = &TreeNode{Val: 1}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{root1}, 7},
		{"test2", args{root2}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob3(tt.args.root); got != tt.want {
				t.Errorf("rob3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"test2", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"test2", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"test3", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 3, 5, 0, 0, 3, 1, 4}}, 6},
		{"test2", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"test3", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"test4", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit4(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2, []int{2, 4, 1}}, 2},
		{"test2", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit4(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit5(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"test2", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit6(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
		{"test2", args{[]int{1, 3, 7, 5, 10, 3}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"test2", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"test3", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 4, 7}}, 3},
		{"test2", args{[]int{2, 2, 2, 2, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
		{"test2", args{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcde", "ace"}, 3},
		{"test2", args{"abc", "abc"}, 3},
		{"test3", args{"abc", "def"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 4, 2}, []int{1, 2, 4}}, 2},
		{"test2", args{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}}, 3},
		{"test3", args{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"test2", args{[]int{1}}, 1},
		{"test3", args{[]int{5, 4, -1, 7, 8}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"abc", "ahbgdc"}, true},
		{"test2", args{"axc", "ahbgdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
