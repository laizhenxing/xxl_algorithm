package bst

import (
	"algo/binarytree"
)

var index, res int = 0, 0

// 二叉搜索树中第K小的元素 [leetcode 230]
func KthSmallest(root *binarytree.TreeNode, k int) int {
	kth_traverse(root, k)
	return res
}

// 中序遍历，升序
func kth_traverse(root *binarytree.TreeNode, k int) {
	if root == nil {
		return
	}

	kth_traverse(root.Left, k)

	index++
	if index == k {
		res = root.Val
		return
	}

	kth_traverse(root.Right, k)
}
