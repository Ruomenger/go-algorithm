package btree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(root *TreeNode, nums *[]int) *[]int {
	if root == nil {
		return nums
	}
	*nums = append(*nums, root.Val)
	if root.Left != nil {
		preorder(root.Left, nums)
	}
	if root.Right != nil {
		preorder(root.Right, nums)
	}
	return nums
}

// preorderTraversalRec 递归版 144. 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal
func preorderTraversalRec(root *TreeNode) []int {
	nums := make([]int, 0)
	preorder(root, &nums)
	return nums
}

// preorderTraversalIter 迭代版 144. 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal
func preorderTraversalIter(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		nums = append(nums, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return nums
}

// inorderTraversalRec 递归版 94. 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversalRec(root *TreeNode) []int {
	nums := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return nums
}

// inorderTraversalIter 迭代版 94. 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal
func inorderTraversalIter(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}
	return nums
}

// postorderTraversal 递归版 145. 二叉树的后序遍历 https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversalRec(root *TreeNode) []int {
	nums := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		nums = append(nums, node.Val)
	}
	postorder(root)
	return nums
}

// postorderTraversalIter 迭代版 145. 二叉树的后序遍历 https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversalIter(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			nums = append(nums, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return nums
}

// levelOrder 102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	nums := make([][]int, 0)
	for queue.Len() != 0 {
		size := queue.Len()
		var tmpArr []int
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		nums = append(nums, tmpArr)
	}
	return nums
}

// levelOrderBottom 107. 二叉树的层序遍历 II https://leetcode.cn/problems/binary-tree-level-order-traversal-ii
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	nums := make([][]int, 0)
	for queue.Len() != 0 {
		size := queue.Len()
		var tmpArr []int
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		nums = append(nums, tmpArr)
	}
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
