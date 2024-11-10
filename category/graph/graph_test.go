package graph

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[][]int{{1, 2}, {3}, {3}, {}}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{"test2", args{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	args1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	args2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	args3 := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{args1}, 1},
		{"test2", args{args2}, 3},
		{"test3", args{args3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
			if got := numIslands2(tt.args.grid); got != tt.want {
				t.Errorf("numIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}
