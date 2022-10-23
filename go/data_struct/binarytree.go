package datastruct

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// 二叉树构造

// 先序遍历（递归）
func PreOrderByRecursive(root *BinaryTree) {
	if root == nil {
		fmt.Println("#")
		return
	}
	// 打印当前节点的值
	fmt.Println(root.Value)
	PreOrderByRecursive(root.Left)
	PreOrderByRecursive(root.Right)
}

// 先序遍历（非递归）
func PreOrderByUnRecursive(root *BinaryTree) {
	if root == nil {
		return
	}

	stack := NewStack()
	stack.Push(root)
	for stack.Len() > 0 {
		node := stack.Pop().(*BinaryTree)
		// 打印当前节点
		fmt.Println(node.Value)
		if node.Right != nil { // 先入右节点
			stack.Push(node.Right)
		}
		if node.Left != nil { // 再入左节点
			stack.Push(node.Left)
		}
	}
}

// 中序遍历（递归）
func InOrderByRecursive(root *BinaryTree) {
	if root == nil {
		return
	}

	InOrderByRecursive(root.Left)
	// 处理当前节点
	fmt.Println(root.Value)
	InOrderByRecursive(root.Right)
}

// 中序遍历（非递归）
func InOrderByUnRecursive(root *BinaryTree) {
	if root == nil {
		return
	}
	stack := NewStack()
	stack.Push(root)
	for !stack.IsEmpty() || root != nil {
		if root != nil {
			stack.Push(root)
			root = root.Left
		} else {
			root = stack.Pop().(*BinaryTree)
			fmt.Println(root.Value)
			root = root.Right
		}
	}
}

// 后序遍历（递归）
func PostOrderByRecursive(root *BinaryTree) {

}

// 后序遍历（非递归）
func PostOrderByUnRecursive(root *BinaryTree) {

}

// 搜索二叉树的判断（BST）

// 完全二叉树的判断

// 宽度优先遍历（队列）

// 满二叉树判断（N(节点) = 2^L(最大深度) -1）

// 平衡二叉树的判断

// 序列化(先序)
func SerializeByPre(root *BinaryTree) string {
	if root == nil {
		return "#_"
	}
	nodeStr := ""
	nodeVal := strconv.Itoa(root.Value)
	nodeStr += nodeVal + "_"

	nodeStr += SerializeByPre(root.Left)
	nodeStr += SerializeByPre(root.Right)
	return nodeStr
}

// 反序列化(先序)	"1_2_4_#_8_#_#_5_9_#_#_#_3_6_#_#_7_#_#_"
func DeserializeByString(nodeStr string) *BinaryTree {
	vals := strings.Split(nodeStr, "_")
	// vals = vals[:len(vals)-1]
	queue := list.New()
	// 存入队列
	for _, v := range vals {
		queue.PushBack(v)
	}
	// 消费队列
	return reconPreOrder(queue)
}

// 先序
func reconPreOrder(queue *list.List) *BinaryTree {
	elem := queue.Front()

	val := (queue.Remove(elem)).(string)

	if val == "#" { // 第一个节点为空
		return nil
	}
	v, _ := strconv.Atoi(val)
	head := &BinaryTree{
		Value: v,
	}
	head.Left = reconPreOrder(queue)
	head.Right = reconPreOrder(queue)
	return head
}

// 凹凸问题
