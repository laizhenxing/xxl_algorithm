package bst

import "algo/binarytree"

var sum int

// BST 转累加树 [leetcode 538]
func ConvertBST(root *binarytree.TreeNode) *binarytree.TreeNode {
	conver_traverse(root)
	return root
}

// 中序遍历，降序
func conver_traverse(root *binarytree.TreeNode) {
	if root == nil {
		return
	}

	conver_traverse(root.Right)

	// 累加右节点
	sum += root.Val
	root.Val = sum

	conver_traverse(root.Left)
}
