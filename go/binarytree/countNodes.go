package binarytree

import (
	"math"
)

// 计算普通二叉树的节点数量
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := CountNodes(root.Left)
	right := CountNodes(root.Right)

	return left + right + 1
}

func CountLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Left == nil && root.Right == nil {
		count += 1
	}

	count += CountLeaves(root.Left)
	count += CountLeaves(root.Right)

	return count
}

// 计算满二叉树
func CountFullBT(root *TreeNode) int {
	var h float64 = 0
	for root != nil {
		root = root.Left
		h++
	}

	return int(math.Pow(2, h)) - 1
}

// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2^h 个节点

// 计算完全二叉树
func CountCompleteBT(root *TreeNode) int {
	l, r := root, root
	var hl, hr float64 = 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hl++
	}
	if hl == hr {
		return int(math.Pow(2, hl)) - 1
	}

	return 1 + CountCompleteBT(root.Left) + CountCompleteBT(root.Right)
}
