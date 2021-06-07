package binarytree

// 待完善

// 前序遍历
func DepthPreoder(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	if root.Left != nil || root.Right != nil {
		depth += 1
	}

	depth = DepthPreoder(root.Left, depth)
	depth = DepthPreoder(root.Right, depth)

	return depth
}

// 后序遍历
func DepthPostoder(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	DepthPreoder(root.Left, depth)
	DepthPreoder(root.Right, depth)

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

	DepthPreoder(root.Left, depth)

	if root.Left != nil || root.Right != nil {
		depth += 1
	}

	DepthPreoder(root.Right, depth)

	return depth
}
