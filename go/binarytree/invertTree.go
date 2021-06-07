package binarytree

// 前序遍历
func InvertTreePreorder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	InvertTreePreorder(root.Left)
	InvertTreePreorder(root.Right)

	return root
}

// 后序遍历
func InvertTreePostorder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	InvertTreePreorder(root.Left)
	InvertTreePreorder(root.Right)

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	return root
}
