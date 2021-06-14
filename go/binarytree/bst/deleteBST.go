package bst

import "algo/binarytree"

// 删除BST中的节点 [leetcode 450]
func DeleteNode(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		// 找到要删除的节点
		// 1. 左右节点都为空，直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 2. 左节点为空，右节点替换删除节点
		if root.Left == nil {
			return root.Right
		}
		// 3. 右节点为空，左节点替换删除节点
		if root.Right == nil {
			return root.Left
		}
		// 4. 左右节点都不为空，2种解法
		// 4.1 找到左子树中最大的节点来替换删除节点
		// 4.2 找到右子树中最小的节点来替换删除节点
		minNode := findMinNode(root.Right)
		// 替换
		root.Val = minNode.Val
		root.Right = DeleteNode(root.Right, minNode.Val)
	} else if root.Val > val {
		root.Left = DeleteNode(root.Left, val)
	} else if root.Val < val {
		root.Right = DeleteNode(root.Right, val)
	}

	return root
}

func findMinNode(node *binarytree.TreeNode) *binarytree.TreeNode {
	if node == nil {
		return node
	}
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func findMaxNode(node *binarytree.TreeNode) *binarytree.TreeNode {
	if node == nil {
		return node
	}
	for node.Right != nil {
		node = node.Right
	}

	return node
}
