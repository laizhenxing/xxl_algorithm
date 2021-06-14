package bst

import "algo/binarytree"

// BST 插入节点 [leetcode 701]
func InsertIntoBST(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil {
		return binarytree.NewTreeNode(val)
	}

	if root.Val < val {
		root.Right = InsertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = InsertIntoBST(root.Left, val)
	}

	return root
}
