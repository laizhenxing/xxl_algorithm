package bst

import "algo/binarytree"

// 判断 BST 的合法性
func IsValidBST(root *binarytree.TreeNode) bool {
	return isValidBST(root, nil, nil)
}

// 限定以 root 为根的子节点必须满足 max.Val > root.Val > min.Val
func isValidBST(root, min, max *binarytree.TreeNode) bool {
	if root == nil {
		return true
	}

	if max != nil && max.Val <= root.Val {
		return false
	}

	if min != nil && min.Val >= root.Val {
		return false
	}

	return isValidBST(root.Left, min, root) && isValidBST(root.Right, root, max)
}
