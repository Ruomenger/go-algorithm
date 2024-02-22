package btree

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
