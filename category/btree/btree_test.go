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

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	root11 := &TreeNode{Val: 1}
	root11.Left = &TreeNode{Val: 3}
	root11.Right = &TreeNode{Val: 2}
	root11.Left.Left = &TreeNode{Val: 5}

	root12 := &TreeNode{Val: 2}
	root12.Left = &TreeNode{Val: 1}
	root12.Right = &TreeNode{Val: 3}
	root12.Left.Right = &TreeNode{Val: 4}
	root12.Right.Right = &TreeNode{Val: 7}

	root13 := &TreeNode{Val: 3}
	root13.Left = &TreeNode{Val: 4}
	root13.Right = &TreeNode{Val: 5}
	root13.Left.Left = &TreeNode{Val: 5}
	root13.Left.Right = &TreeNode{Val: 4}
	root13.Right.Right = &TreeNode{Val: 7}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1: root11, root2: root12}, root13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, 2}, root1.Left},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 6}
	root1.Right.Left = &TreeNode{Val: 3}
	root1.Right.Right = &TreeNode{Val: 7}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 5}
	root2.Left.Left = &TreeNode{Val: 0}
	root2.Left.Right = &TreeNode{Val: 2}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 6}
	root2.Left.Right.Right = &TreeNode{Val: 3}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{root1}, false},
		{"test2", args{root2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{Val: 1}

	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}
	root2.Right.Left = &TreeNode{Val: 2}

	root3 := &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{root1}, []int{1}},
		{"test2", args{root2}, []int{2}},
		{"test3", args{root3}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 0}
	root1.Right.Right = &TreeNode{Val: 8}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	p1 := root1.Left
	q1 := root1.Right

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Right = &TreeNode{Val: 1}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 2}
	root2.Right.Left = &TreeNode{Val: 0}
	root2.Right.Right = &TreeNode{Val: 8}
	root2.Left.Right.Left = &TreeNode{Val: 7}
	root2.Left.Right.Right = &TreeNode{Val: 4}
	p2 := root2.Left
	q2 := root2.Left.Right.Right
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, p1, q1}, root1},
		{"test2", args{root2, p2, q2}, p2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 6}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 0}
	root2.Right = &TreeNode{Val: 48}
	root2.Right.Left = &TreeNode{Val: 12}
	root2.Right.Right = &TreeNode{Val: 49}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{root1}, 1},
		{"test2", args{root1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestor2(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	root1 := &TreeNode{Val: 6}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 8}
	root1.Left.Left = &TreeNode{Val: 0}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Left.Right.Left = &TreeNode{Val: 3}
	root1.Left.Right.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 7}
	root1.Right.Right = &TreeNode{Val: 9}
	p1 := root1.Left
	q1 := root1.Right

	root2 := &TreeNode{Val: 6}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 8}
	root2.Left.Left = &TreeNode{Val: 0}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Left.Right.Left = &TreeNode{Val: 3}
	root2.Left.Right.Right = &TreeNode{Val: 5}
	root2.Right.Left = &TreeNode{Val: 7}
	root2.Right.Right = &TreeNode{Val: 9}
	p2 := root2.Left
	q2 := root2.Left.Right

	root3 := &TreeNode{Val: 2}
	root3.Left = &TreeNode{Val: 1}
	root3.Right = &TreeNode{Val: 3}
	p3 := root3.Right
	q3 := root3.Left
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, p1, q1}, root1},
		{"test2", args{root2, p2, q2}, p2},
		{"test3", args{root3, p3, q3}, root3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}

	ans1 := &TreeNode{Val: 4}
	ans1.Left = &TreeNode{Val: 2}
	ans1.Right = &TreeNode{Val: 7}
	ans1.Right.Left = &TreeNode{Val: 5}
	ans1.Left.Left = &TreeNode{Val: 1}
	ans1.Left.Right = &TreeNode{Val: 3}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, 5}, ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	root1 := &TreeNode{Val: 5}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 6}
	root1.Left.Left = &TreeNode{Val: 2}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Right.Right = &TreeNode{Val: 7}

	want1 := &TreeNode{Val: 5}
	want1.Left = &TreeNode{Val: 4}
	want1.Right = &TreeNode{Val: 6}
	want1.Left.Left = &TreeNode{Val: 2}
	want1.Right.Right = &TreeNode{Val: 7}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, 3}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 0}
	root1.Right = &TreeNode{Val: 2}
	want1 := &TreeNode{Val: 1}
	want1.Right = &TreeNode{Val: 2}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 0}
	root2.Right = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 2}
	root2.Left.Right.Left = &TreeNode{Val: 1}
	want2 := &TreeNode{Val: 3}
	want2.Left = &TreeNode{Val: 2}
	want2.Left.Left = &TreeNode{Val: 1}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{root1, 1, 2}, want1},
		{"test2", args{root2, 1, 3}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	want1 := &TreeNode{Val: 5}
	want1.Left = &TreeNode{Val: 3}
	want1.Right = &TreeNode{Val: 8}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"test1", args{[]int{3, 5, 8}}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
