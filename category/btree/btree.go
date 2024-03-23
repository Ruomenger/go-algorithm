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

// rightSideView 199. 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := list.New()
	queue.PushBack(root)
	nums := make([]int, 0)
	for queue.Len() != 0 {
		size := queue.Len()
		var tempNode *TreeNode
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == size-1 {
				tempNode = node
			}
		}
		nums = append(nums, tempNode.Val)
	}
	return nums
}

// invertTree 226. 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

// isSymmetricRec 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree
func isSymmetricRec(root *TreeNode) bool {
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root, root)
}

// isSymmetricIter 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree
func isSymmetricIter(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(root)
	for queue.Len() != 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}

// maxDepth 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// minDepth 111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return min(minDepth(root.Left)+1, minDepth(root.Right)+1)
}

// minDepth2 111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	dep := 0
	for queue.Len() != 0 {
		dep++
		// TODO 这里注意要把长度保存下来，否则push数据后会影响长度
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return dep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return -1
	}
	return max(leftDepth+1, rightDepth+1)
}

// isBalanced 110. 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree
func isBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}

// countNodes 222. 完全二叉树的节点个数 https://leetcode.cn/problems/count-complete-tree-nodes
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	for left := root.Left; left != nil; left = left.Left {
		leftHeight++
	}
	for right := root.Right; right != nil; right = right.Right {
		rightHeight++
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
