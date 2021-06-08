package binarytree

// 填充二叉树节点的右侧指针 [leetcode 114]
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 1. 把左右子树展平
	Flatten(root.Left)
	Flatten(root.Right)

	left := root.Left
	right := root.Right

	// 2. 把左子树置空，把左节点挂到右子树
	root.Left = nil
	root.Right = left

	// 找到最后一个右节点
	p := root
	for p.Right != nil {
		p = p.Right
	}
	// 3. 把之前有节点接到当前右子树上
	p.Right = right
}
