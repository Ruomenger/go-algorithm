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
