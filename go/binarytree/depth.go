package binarytree

// 待完善

// 前序遍历
func DepthPreoder(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(DepthPreoder(root.Left), DepthPreoder(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 后序遍历
func DepthPostoder(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	DepthPostoder(root.Left, depth)
	DepthPostoder(root.Right, depth)

	if root.Left != nil || root.Right != nil {
		depth += 1
	}

	return depth
}

// 中序遍历
func DepthInorder(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	DepthInorder(root.Left, depth)

	if root.Left != nil || root.Right != nil {
		depth += 1
	}

	DepthInorder(root.Right, depth)

	return depth
}
