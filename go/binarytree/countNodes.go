package binarytree

// 计算二叉树的节点数量
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
