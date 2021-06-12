package binarytree

import "math"

// 构建最大二叉树 [leetcode 654]
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return buildNode(nums, 0, len(nums)-1)
}

func buildNode(nums []int, low, hight int) *TreeNode {
	if low > hight {
		return nil
	}

	index := -1
	maxVal := math.MinInt64

	for i := low; i <= hight; i++ {
		if nums[i] > maxVal {
			index = i
			maxVal = nums[i]
		}
	}

	root := &TreeNode{Val: maxVal}
	root.Left = buildNode(nums, low, index-1)
	root.Right = buildNode(nums, index+1, hight)

	return root
}
