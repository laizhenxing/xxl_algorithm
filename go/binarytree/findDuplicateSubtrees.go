package binarytree

import "fmt"

var res []*TreeNode
var nodes map[string]int

// 寻找重复子树 [leetcode 652]
func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res = make([]*TreeNode, 0)
	nodes = make(map[string]int, 0)

	traverse(root)
	return res
}

// 辅助函数
func traverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	// 左子树
	left := traverse(root.Left)
	// 右子树
	right := traverse(root.Right)

	// 后序遍历代码
	// 序列化
	node := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
	// 是否存在，存在则将当前节点加入结果集
	if num, ok := nodes[node]; ok && num == 1 {
		res = append(res, root)
	}

	nodes[node] += 1

	return node
}
