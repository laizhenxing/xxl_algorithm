package binarytree

type PerfictTreeNode struct {
	Val   int
	Left  *PerfictTreeNode
	Right *PerfictTreeNode
	Next  *PerfictTreeNode
}

// 填充二叉树节点的右侧指针 [leetcode 116]
func Connect(root *PerfictTreeNode) *PerfictTreeNode {
	if root == nil {
		return nil
	}
	connect(root.Left, root.Right)
	return root
}

func connect(node1, node2 *PerfictTreeNode) {
	if node1 == nil || node2 == nil {
		return
	}

	// 将传入的两个节点连接
	node1.Next = node2

	// 连接节点1
	connect(node1.Left, node1.Right)
	// 连接节点2
	connect(node2.Left, node2.Right)

	// 连接节点1，2
	connect(node1.Right, node2.Left)
}
