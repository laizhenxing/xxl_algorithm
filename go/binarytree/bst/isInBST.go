package bst

import "algo/binarytree"

func IsInBST(root *binarytree.TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}
	if root.Val > target {
		return IsInBST(root.Left, target)
	}
	if root.Val < target {
		return IsInBST(root.Right, target)
	}
	return false
}
