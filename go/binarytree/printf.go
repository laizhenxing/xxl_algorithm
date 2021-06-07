package binarytree

import "fmt"

var nodes [][]int

// 待完善
func Printf(root *TreeNode) {
	if root == nil {
		return
	}

	nodes = [][]int{
		[]int{root.Val},
	}

	traverse(root, 0)

	fmt.Println(nodes)
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}

	fmt.Println("root.val: ", root.Val, len(nodes))
	if len(nodes) <= k {
		subNode := []int{}
		nodes = append(nodes, subNode)
	}

	nodes[k] = append(nodes[k], root.Val)

	traverse(root.Left, k+1)
	traverse(root.Right, k+1)
}
