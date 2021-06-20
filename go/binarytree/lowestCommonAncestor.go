package binarytree

import "reflect"

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if reflect.DeepEqual(root, p) || reflect.DeepEqual(root, q) {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 情况1: p,q 在 root 中，left,right 分别就是 p, q
	if left != nil && right != nil {
		return root
	}
	// 情况2: p, q 都不在 root 中，返回 nil
	if left == nil && right == nil {
		return nil
	}
	// 情况3: p, q 只有一个存在 root 中，返回存在的那个节点
	if left == nil {
		return right
	}
	return left
}
