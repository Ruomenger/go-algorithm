package btree

import (
	"container/list"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func getPath(root *TreeNode, prefix string, ans []string) []string {
	if root == nil {
		return ans
	}
	if prefix != "" {
		prefix = prefix + "->"
	}
	if root.Left == nil && root.Right == nil {
		ans = append(ans, prefix+strconv.Itoa(root.Val))
		return ans
	}
	if root.Left != nil {
		ans = getPath(root.Left, prefix+strconv.Itoa(root.Val), ans)
	}
	if root.Right != nil {
		ans = getPath(root.Right, prefix+strconv.Itoa(root.Val), ans)
	}
	return ans
}

// binaryTreePaths 257. 二叉树的所有路径 https://leetcode.cn/problems/binary-tree-paths
func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	return getPath(root, "", ans)
}

// sumOfLeftLeaves 404. 左叶子之和 https://leetcode.cn/problems/sum-of-left-leaves
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	sum := 0
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					sum += node.Left.Val
				} else {
					queue.PushBack(node.Left)
				}
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return sum
}

// findBottomLeftValue 513. 找树左下角的值 https://leetcode.cn/problems/find-bottom-left-tree-value
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return ans
}

// hasPathSum 112. 路径总和 https://leetcode.cn/problems/path-sum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
		return false
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}
	inorderMap := make(map[int]int, len(inorder))
	for i, num := range inorder {
		inorderMap[num] = i
	}
	rootIdx := inorderMap[postorder[length-1]]
	root := &TreeNode{
		Val: postorder[length-1],
	}
	var buildTreeFrom func(inStart, inEnd, postStart, postEnd int) *TreeNode
	buildTreeFrom = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		idx := inorderMap[postorder[postEnd]]
		node := &TreeNode{
			Val: postorder[postEnd],
		}

		node.Left = buildTreeFrom(inStart, idx-1, postStart, postStart+idx-1-inStart)
		node.Right = buildTreeFrom(idx+1, inEnd, postEnd+idx-inEnd, postEnd-1)
		return node
	}
	root.Left = buildTreeFrom(0, rootIdx-1, 0, rootIdx-1)
	root.Right = buildTreeFrom(rootIdx+1, length-1, rootIdx, length-2)
	return root
}

// constructMaximumBinaryTree 654. 最大二叉树 https://leetcode.cn/problems/maximum-binary-tree
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	num2IdxMap := make(map[int]int, len(nums))
	maxNum := math.MinInt
	for i, num := range nums {
		num2IdxMap[num] = i
		if num > maxNum {
			maxNum = num
		}
	}
	idx := num2IdxMap[maxNum]
	// idx起始可以省略，因为自己可以找到
	var buildTreeFrom func(start, idx, end int) *TreeNode
	buildTreeFrom = func(start, idx, end int) *TreeNode {
		if start > end {
			return nil
		}
		node := &TreeNode{Val: nums[idx]}
		leftMaxNum := math.MinInt
		for i := start; i < idx; i++ {
			if nums[i] > leftMaxNum {
				leftMaxNum = nums[i]
			}
		}
		node.Left = buildTreeFrom(start, num2IdxMap[leftMaxNum], idx-1)
		rightMaxNum := math.MinInt
		for i := idx + 1; i <= end; i++ {
			if nums[i] > rightMaxNum {
				rightMaxNum = nums[i]
			}
		}
		node.Right = buildTreeFrom(idx+1, num2IdxMap[rightMaxNum], end)
		return node
	}
	return buildTreeFrom(0, idx, len(nums)-1)
}

// mergeTrees 617. 合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	node := &TreeNode{Val: root1.Val + root2.Val}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	return node
}

// searchBST 700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

// isValidBST 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, minVal, maxVal int) bool
	isValid = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return true
		}
		if root.Left != nil && !(root.Left.Val < root.Val && root.Left.Val > minVal) {
			return false
		}
		if root.Right != nil && !(root.Right.Val > root.Val && root.Right.Val < maxVal) {
			return false
		}
		return isValid(root.Left, minVal, root.Val) && isValid(root.Right, root.Val, maxVal)
	}

	return isValid(root, math.MinInt, math.MaxInt)
}

// findMode 501. 二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	count := 0
	maxCount := 0
	var pre *TreeNode
	var find func(cur *TreeNode)
	find = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		find(cur.Left)
		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		pre = cur
		if count == maxCount {
			ret = append(ret, cur.Val)
		} else if count > maxCount {
			maxCount = count
			ret = []int{cur.Val}
		}
		find(cur.Right)
	}
	find(root)
	return ret
}

// lowestCommonAncestor 236. 二叉树的最近公共祖先 medium
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	if root.Left == nil {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Right == nil {
		return lowestCommonAncestor(root.Left, p, q)
	}
	leftRoot := lowestCommonAncestor(root.Left, p, q)
	if leftRoot == nil {
		return lowestCommonAncestor(root.Right, p, q)
	}
	rightRoot := lowestCommonAncestor(root.Right, p, q)
	if rightRoot == nil {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func getMinimumDifference(root *TreeNode) int {
	minVal := math.MaxInt
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var traversal func(root *TreeNode)
	var prev *TreeNode
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if prev != nil {
			minVal = min(node.Val-prev.Val, minVal)
		}
		prev = node
		traversal(node.Right)
	}
	traversal(root)
	return minVal
}

// lowestCommonAncestor2 235. 二叉搜索树的最近公共祖先 medium
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val < q.Val {
		return lowestCommonAncestorBST(root, p, q)
	}
	return lowestCommonAncestorBST(root, q, p)
}

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorBST(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val > root.Val {
		return root
	}
	return nil
}

// insertIntoBST 701. 二叉搜索树中的插入操作 medium
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
		return root
	}
	root.Right = insertIntoBST(root.Right, val)
	return root
}

// deleteNode 450. 删除二叉搜索树中的节点 medium
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil && root.Right == nil {
		return nil
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	node := root.Right
	for ; node.Left != nil; node = node.Left {
	}
	node.Left = root.Left
	return root.Right
}

// trimBST 669. 修剪二叉搜索树 medium
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// sortedArrayToBST 108. 将有序数组转换为二叉搜索树 simple
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[0:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}

// convertBST 538. 把二叉搜索树转换为累加树 meduim
func convertBST(root *TreeNode) *TreeNode {
	var build func(node *TreeNode, sum int) int
	build = func(node *TreeNode, sum int) int {
		if node == nil {
			return sum
		}
		node.Val += build(node.Right, sum)
		return build(node.Left, node.Val)
	}
	build(root, 0)
	return root
}

// sortedListToBST 109. 有序链表转换二叉搜索树 medium
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}

// removeLeafNodes 1325. 删除给定值的叶子节点
// tag: 二叉树
// level: medium
// https://leetcode.cn/problems/delete-leaves-with-a-given-value/description/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	var delete func(*TreeNode) (*TreeNode, bool)
	delete = func(node *TreeNode) (*TreeNode, bool) {
		if node == nil {
			return nil, true
		}
		leftNode, leftIs := delete(node.Left)
		rightNode, rightIs := delete(node.Right)
		if leftIs && leftNode != nil && leftNode.Val == target {
			leftNode = nil
		}
		if rightIs && rightNode != nil && rightNode.Val == target {
			rightNode = nil
		}

		node.Left = leftNode
		node.Right = rightNode
		if leftNode == nil && rightNode == nil {
			return node, true
		}
		return node, false
	}

	node, is := delete(root)
	if is && node.Val == target {
		return nil
	}
	return node
}
