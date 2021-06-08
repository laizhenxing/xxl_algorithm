package binarytree

// 二叉树的问题难点在于，如何把题目的要求细化成每个节点需要做的事情
// 二叉树节点
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// N 叉树节点
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// 前序遍历框架
func traverse_front(root *TreeNode) {
	if root == nil {
		return
	}

	/**********前序遍历位置*************/
	// TODO
	/**********前序遍历位置*************/

	traverse_front(root.Left)
	traverse_front(root.Right)
}
