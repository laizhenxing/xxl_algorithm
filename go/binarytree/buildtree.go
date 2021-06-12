package binarytree

func BuildTreeFromPreAndIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	index := 0
	for i, v := range inorder {
		if v == preorder[0] {
			index = i
			break
		}
	}

	left_in := inorder[:index]
	right_in := inorder[index+1:]

	left_pre := preorder[1:(1 + len(left_in))]
	right_pre := preorder[len(preorder)-len(right_in):]

	root.Left = BuildTreeFromPreAndIn(left_pre, left_in)
	root.Right = BuildTreeFromPreAndIn(right_pre, right_in)

	return root
}

func BuildTreeFromInAndPost(inorder, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	postLen := len(postorder)
	rootVal := postorder[postLen-1]
	root := &TreeNode{Val: rootVal}
	index := 0
	for i, v := range inorder {
		if v == rootVal {
			index = i
			break
		}
	}

	left_in := inorder[0:index]
	left_post := postorder[0:index]

	right_in := inorder[index+1:]
	right_post := postorder[index:(index + len(right_in))]

	root.Left = BuildTreeFromInAndPost(left_in, left_post)
	root.Right = BuildTreeFromInAndPost(right_in, right_post)

	return root
}
