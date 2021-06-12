package binarytree

import "fmt"

var _nodes [][]int

// 待完善
func Printf(root *TreeNode) {
	if root == nil {
		return
	}

	_nodes = [][]int{
		[]int{root.Val},
	}

	traverse_print(root, 0)

	fmt.Println(nodes)
}

func traverse_print(root *TreeNode, k int) {
	if root == nil {
		return
	}

	fmt.Println("root.val: ", root.Val, len(nodes))
	if len(nodes) <= k {
		subNode := []int{}
		_nodes = append(_nodes, subNode)
	}

	_nodes[k] = append(_nodes[k], root.Val)

	traverse_print(root.Left, k+1)
	traverse_print(root.Right, k+1)
}
