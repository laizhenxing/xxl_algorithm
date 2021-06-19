package bst

import (
	"algo/binarytree"
	"math"
)

var maxSum = 0

func MaxSumBST(root *binarytree.TreeNode) int {
	traverse_maxSumBST(root)
	return maxSum
}

// res[0]: 以 root 为根的二叉树是否是 BST， 1-是，0-不是
// res[1]: 以 root 为根的二叉树所有节点中的最小值
// res[2]: 以 root 为根的二叉树所有节点中的最大值
// res[3]: 以 root 为根的二叉树所有节点之和
func traverse_maxSumBST(root *binarytree.TreeNode) []int {
	// base case
	if root == nil {
		return []int{1, math.MinInt64, math.MaxInt64, 0}
	}

	// 计算左右子树
	left := traverse_maxSumBST(root.Left)
	right := traverse_maxSumBST(root.Right)

	// 后序遍历
	res := make([]int, 4)
	// 判断以 root 为根的二叉树是不是 BST
	if left[0] == 1 && right[0] == 1 &&
		root.Val > left[2] && root.Val < right[1] {
		// 以 root 为根的二叉树是 BST
		res[0] = 1
		// 计算以 root 为根的这棵树 BST 的最小值
		res[1] = min(left[1], root.Val)
		// 计算以 root 为根的这棵树 BST 的最大值
		res[2] = max(right[2], root.Val)
		// 计算以 root 为根的这棵 BST 所有节点之和
		res[3] = left[3] + right[3] + root.Val
		// 更新全局变量
		maxSum = max(res[3], maxSum)
	} else {
		res[0] = 0
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
