package btree

import (
	"reflect"
	"testing"
)

func TestPreorderTraversalRec(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Right = &TreeNode{Val: 2}
	test1.root.Right.Left = &TreeNode{Val: 3}
	test2 := args{}
	test3 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test4 := args{
		&TreeNode{
			Val: 3,
		},
	}
	test4.root.Left = &TreeNode{Val: 1}
	test4.root.Right = &TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", test1, []int{1, 2, 3}},
		{"test2", test2, []int{}},
		{"test3", test3, []int{1}},
		{"test4", test4, []int{3, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalRec(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalRec() = %v, want %v", got, tt.want)
			}
			if got := preorderTraversalIter(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Right = &TreeNode{Val: 2}
	test1.root.Right.Right = &TreeNode{Val: 3}
	test2 := args{}
	test3 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test4 := args{
		&TreeNode{
			Val: 3,
		},
	}
	test4.root.Left = &TreeNode{Val: 1}
	test4.root.Right = &TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", test1, []int{1, 2, 3}},
		{"test2", test2, []int{}},
		{"test3", test3, []int{1}},
		{"test4", test4, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalRec(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalRec() = %v, want %v", got, tt.want)
			}
			if got := inorderTraversalIter(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostorderTraversalRec(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Right = &TreeNode{Val: 2}
	test1.root.Right.Right = &TreeNode{Val: 3}
	test2 := args{}
	test3 := args{
		&TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}
	test4 := args{
		&TreeNode{
			Val: 3,
		},
	}
	test4.root.Left = &TreeNode{Val: 1}
	test4.root.Right = &TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", test1, []int{3, 2, 1}},
		{"test2", test2, []int{}},
		{"test3", test3, []int{1}},
		{"test4", test4, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversalRec(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalRec() = %v, want %v", got, tt.want)
			}
			if got := postorderTraversalIter(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Left = &TreeNode{Val: 9}
	test1.root.Right = &TreeNode{Val: 20}
	test1.root.Right.Left = &TreeNode{Val: 15}
	test1.root.Right.Right = &TreeNode{Val: 7}
	test2 := args{
		&TreeNode{
			Val: 1,
		},
	}
	test3 := args{}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", test1, [][]int{{3}, {9, 20}, {15, 7}}},
		{"test2", test2, [][]int{{1}}},
		{"test3", test3, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Left = &TreeNode{Val: 9}
	test1.root.Right = &TreeNode{Val: 20}
	test1.root.Right.Left = &TreeNode{Val: 15}
	test1.root.Right.Right = &TreeNode{Val: 7}
	test2 := args{
		&TreeNode{
			Val: 1,
		},
	}
	test3 := args{}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", test1, [][]int{{15, 7}, {9, 20}, {3}}},
		{"test2", test2, [][]int{{1}}},
		{"test3", test3, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}}
	test1.root.Left = &TreeNode{Val: 9}
	test1.root.Right = &TreeNode{Val: 20}
	test1.root.Right.Left = &TreeNode{Val: 15}
	test1.root.Right.Right = &TreeNode{Val: 7}
	test2 := args{
		&TreeNode{
			Val: 1,
		},
	}
	test3 := args{}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", test1, []int{3, 20, 7}},
		{"test2", test2, []int{1}},
		{"test3", test3, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
