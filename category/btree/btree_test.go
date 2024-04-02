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

func TestIsSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	test1 := args{
		&TreeNode{
			Val: 1,
		}}
	test1.root.Left = &TreeNode{Val: 2}
	test1.root.Right = &TreeNode{Val: 2}
	test1.root.Left.Left = &TreeNode{Val: 3}
	test1.root.Left.Right = &TreeNode{Val: 4}
	test1.root.Right.Left = &TreeNode{Val: 4}
	test1.root.Right.Right = &TreeNode{Val: 3}
	test2 := args{
		&TreeNode{
			Val: 1,
		},
	}
	test2.root.Left = &TreeNode{Val: 2}
	test2.root.Right = &TreeNode{Val: 2}
	test2.root.Left.Right = &TreeNode{Val: 3}
	test2.root.Right.Right = &TreeNode{Val: 3}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", test1, true},
		{"test2", test2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricRec(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isSymmetricRec() = %v, want %v", got, tt.want)
			}
			if got := isSymmetricIter(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isSymmetricIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}
	root2.Right.Right = &TreeNode{Val: 3}
	root2.Right.Right.Right = &TreeNode{Val: 4}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{root1}, 2},
		{"test2", args{root2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
			if got := minDepth2(tt.args.root); got != tt.want {
				t.Errorf("minDepth2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Left.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Right = &TreeNode{Val: 4}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test2", args{root2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	root1 := &TreeNode{
		Val: 3,
	}
	root1.Left = &TreeNode{
		Val: 9,
	}
	root1.Right = &TreeNode{
		Val: 20,
	}
	root1.Right.Left = &TreeNode{
		Val: 15,
	}
	root1.Right.Right = &TreeNode{
		Val: 7,
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}}, root1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	root1 := &TreeNode{Val: 6}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 5}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Right = &TreeNode{Val: 1}
	root1.Right.Left = &TreeNode{Val: 0}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{[]int{3, 2, 1, 6, 0, 5}}, root1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
