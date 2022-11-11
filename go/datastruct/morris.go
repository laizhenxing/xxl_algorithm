package datastruct

import (
	"fmt"
)

// morris 遍历
// 细节：
// 假设来到当前节点 cur, 开始 cur 从头节点开始
// 1. 如果 cur 没有左孩子，cur 向右移动(cur=cur.Right)
// 2. 如果 cur 有左孩子，找到左子树上最右的节点 mostRight:
// 	a) 如果 mostRight 的右指针指向空，让其指向 cur (mostRight.Right = cur),然后 cur 向左移动（cur=cur.Left）
// 	b) 如果 mostRight 的右节点指向 cur, 让其指向 null (mostRight.Right=null), 然后 cur 向右移动（cur=cur.Right）
// 3. cur==null 遍历停止

func Morris(root *BinaryTree) {
	if root == nil {
		return
	}
	cur := root // cur == head
	var mostRight *BinaryTree = nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil { // cur 左节点不为空
			for mostRight.Right != nil &&
				mostRight.Right != cur { // 找到左节点的最右节点
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil { // 最右节点的右节点为空，第一次到达该节点
				mostRight.Right = cur // 将最右节点右节点指向 cur
				cur = cur.Left        // cur 移动到左节点
				continue
			} else {
				mostRight.Right = nil // 第二次到达该节点，将该节点的右节点恢复为 nil，cur 向右移动
			}
		}
		cur = cur.Right // cur 向右移动
	}
}

// moriis 前序遍历
func MorrisPreorder(root *BinaryTree) {
	if root == nil {
		return
	}
	cur := root
	var mostRight *BinaryTree = nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil { // 第一次到达该节点
				mostRight.Right = cur
				cur = cur.Left
				// 前序
				fmt.Println("前序遍历!!!")
				continue
			} else {
				mostRight.Right = nil
			}
		} else { // 没有左树
			// 前序
			fmt.Println("前序遍历!!!")
		}
		cur = cur.Right
	}
}

// moriis 中序遍历
func MorrisInorder(root *BinaryTree) {
	var cur, mostRight *BinaryTree = root, nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
				// 中序遍历
				fmt.Println("中序遍历")
			}
		} else {
			// 中序遍历
			fmt.Println("中序遍历")
		}
		cur = cur.Right
	}
}

// moriis 后序遍历
func MorrisPostorder(root *BinaryTree) {
	var cur, mostRight *BinaryTree = root, nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
				printEdge(cur.Left)
			}
		}
		cur = cur.Right
	}
	printEdge(root)
}

// 逆序打印这棵树的右边界
func printEdge(node *BinaryTree) {
	tail := reverseEdge(node)
	cur := tail
	for cur != nil {
		fmt.Println(cur.Value)
		cur = cur.Right
	}
	reverseEdge(tail)
}

func reverseEdge(node *BinaryTree) *BinaryTree {
	var pre, next *BinaryTree = nil, nil
	for node != nil {
		next = node.Right
		node.Right = pre
		pre = node
		node = next
	}
	return pre
}
