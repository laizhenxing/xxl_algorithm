package datastruct

import (
	"container/list"
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
func PreOrderByRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	// 打印当前节点的值
	vals = append(vals, root.Value)
	vleft := PreOrderByRecursive(root.Left)
	vright := PreOrderByRecursive(root.Right)
	vals = append(vals, vleft...)
	vals = append(vals, vright...)
	return
}

// 先序遍历（非递归）
func PreOrderByUnRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}

	stack := NewStack()
	stack.Push(root)
	for stack.Len() > 0 {
		node := stack.Pop().(*BinaryTree)
		// 打印当前节点
		vals = append(vals, node.Value)
		if node.Right != nil { // 先入右节点
			stack.Push(node.Right)
		}
		if node.Left != nil { // 再入左节点
			stack.Push(node.Left)
		}
	}
	return
}

// 中序遍历（递归）
func InOrderByRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}

	vleft := InOrderByRecursive(root.Left)
	vals = append(vals, vleft...)
	// 处理当前节点
	vals = append(vals, root.Value)
	vright := InOrderByRecursive(root.Right)
	vals = append(vals, vright...)
	return
}

// 中序遍历（非递归）
// 1. 申请一个栈stack，记住cur=root
// 2. 先把 cur 压入栈，cur=cur.left 重复2
// 3. 知道 cur==nil, 弹出 stack 一个节点 node，让 cur=node.right, 然后重复步骤2
// 4. 当栈为空且cur为空时，整个过程停止
func InOrderByUnRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	stack := NewStack()
	for !stack.IsEmpty() || root != nil {
		if root != nil {
			stack.Push(root) // root 即为cur, 入栈
			root = root.Left // 在转向左节点
		} else { // 当 cur 为空时，弹栈
			root = stack.Pop().(*BinaryTree)
			vals = append(vals, root.Value)
			root = root.Right // cur 指向弹出节点的右子树
		}
	}
	return
}

// 后序遍历（递归）
func PostOrderByRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	vleft := PostOrderByRecursive(root.Left)
	vright := PostOrderByRecursive(root.Right)
	// 处理当前节点
	vals = append(vals, vleft...)
	vals = append(vals, vright...)
	vals = append(vals, root.Value)
	return
}

// 后序遍历（非递归）
// 采用双栈实现
// 1. 当前节点入栈1，栈1不为空就弹出一个节点，将左节点，右节点依次压入栈1
// 2. 栈1 弹出的节点放入栈2
// 3. 重复1，2.
// 4. 最后依次弹出栈2的节点，即为后续遍历的节点
// 边界：栈1 不为空
func PostOrderByUnRecursive(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	s1, s2 := NewStack(), NewStack()
	s1.Push(root)
	for !s1.IsEmpty() {
		node := s1.Pop().(*BinaryTree)
		s2.Push(node) // 弹出的节点压入s2
		if node.Left != nil {
			s1.Push(node.Left)
		}
		if node.Right != nil {
			s1.Push(node.Right)
		}
	}
	// 输出结果
	for !s2.IsEmpty() {
		node := s2.Pop().(*BinaryTree)
		vals = append(vals, node.Value)
	}
	return
}

type (
	// BST 节点信息
	NodeBSTData struct {
		min   int // 右边最小值
		max   int // 左边的最大值
		isBST bool
	}

	// 满二叉树节点信息
	NodeFullTreeData struct {
		height int // 满二叉树的最深高度
		nodes  int // 所有节点树
	}

	//平衡二叉树节点信息
	NodeBalanceData struct {
		height    int
		isBalance bool
	}
)

func NewNodeBSTData(min, max int, isBST bool) *NodeBSTData {
	return &NodeBSTData{
		min:   min,
		max:   max,
		isBST: isBST,
	}
}

func (bst *NodeBSTData) IsBST() bool {
	return bst.isBST
}

func (bst *NodeBSTData) Min() int {
	return bst.min
}

func (bst *NodeBSTData) Max() int {
	return bst.max
}

func NewNodeFullTreeData(height, nodes int) *NodeFullTreeData {
	return &NodeFullTreeData{
		height: height,
		nodes:  nodes,
	}
}

func (nf *NodeFullTreeData) Height() int {
	return nf.height
}

func (nf *NodeFullTreeData) Nodes() int {
	return nf.nodes
}

func NewNodeBalanceData(height int, isBalance bool) *NodeBalanceData {
	return &NodeBalanceData{
		height:    height,
		isBalance: isBalance,
	}
}

func (nb *NodeBalanceData) Height() int {
	return nb.height
}

func (nb *NodeBalanceData) IsBalance() bool {
	return nb.isBalance
}

// 搜索二叉树的判断（BST）
// 特点：每棵子树都满足： 左 < 头 < 右
func IsBST(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	NodeBSTData := GetNodeBSTData(root)
	return NodeBSTData.IsBST()
}

func IsBST2(root *BinaryTree) bool {
	if root != nil {
		stack := NewStack()
		preValue := INT_MIN

		for !stack.IsEmpty() || root != nil {
			if root != nil {
				stack.Push(root)
				root = root.Left
			} else {
				root = stack.Pop().(*BinaryTree)
				if root.Value <= preValue {
					return false
				}
				preValue = root.Value
				root = root.Right
			}
		}
	}
	return true
}

var preValue int

// 判断是否是二叉搜索树
func IsBST3(root *BinaryTree) bool {
	preValue = INT_MIN
	return isBST3(root)
}

func isBST3(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	bLeft := isBST3(root.Left)
	if !bLeft {
		return false
	}
	if root.Value <= preValue {
		return false
	}
	preValue = root.Value
	return isBST3(root.Right)
}

// 获取当前节点的信息
func GetNodeBSTData(node *BinaryTree) *NodeBSTData {
	if node == nil {
		return NewNodeBSTData(0, 0, true)
	}
	leftData := GetNodeBSTData(node.Left)
	rightData := GetNodeBSTData(node.Right)

	min, max, isBST := node.Value, node.Value, true
	if node.Left != nil {
		max = getMax(max, leftData.Max())
		min = getMin(min, leftData.Min())
	}
	if node.Right != nil {
		max = getMax(max, rightData.Max())
		min = getMin(min, rightData.Min())
	}

	// 左节点不为空，且（左节点不是 BST或者当前节点的值<=左子树的最大值
	if (node.Left != nil) && (!leftData.IsBST() || leftData.Max() >= node.Value) {
		isBST = false
	}
	// 右节点不为空，且（右节点不是 BST或者当前节点的值>=右子树的最小值
	if (node.Right != nil) && (!rightData.IsBST() || rightData.Min() <= node.Value) {
		isBST = false
	}

	return NewNodeBSTData(min, max, isBST)
}

// 满二叉树判断（N(节点) = 2^L(最大深度) -1）
func IsFullTree(node *BinaryTree) bool {
	if node == nil {
		return true
	}
	isFull := true
	fullData := GetNodeFullData(node)
	if fullData.Nodes() != (1<<fullData.Height() - 1) {
		isFull = false
	}
	return isFull
}

// 获取满二叉树节点信息
func GetNodeFullData(node *BinaryTree) *NodeFullTreeData {
	if node == nil {
		return NewNodeFullTreeData(0, 0)
	}
	leftData := GetNodeFullData(node.Left)
	rightData := GetNodeFullData(node.Right)

	height := getMax(leftData.height, rightData.height) + 1
	nodes := leftData.nodes + rightData.nodes + 1

	return NewNodeFullTreeData(height, nodes)
}

// 平衡二叉树的判断，满足以下3个条件
// 1. 左树是平衡树
// 2. 右树是平衡树
// 3. 左右子树高度差<=1
func IsBalanceTree(node *BinaryTree) bool {
	if node == nil {
		return true
	}
	nodeData := GetNodeBalanceData(node)
	return nodeData.IsBalance()
}

// 获取平衡二叉树节点信息
func GetNodeBalanceData(node *BinaryTree) *NodeBalanceData {
	if node == nil {
		return NewNodeBalanceData(0, true)
	}

	leftData := GetNodeBalanceData(node.Left)
	rightData := GetNodeBalanceData(node.Right)

	height := getMax(leftData.Height(), rightData.Height()) + 1
	isBalance := leftData.IsBalance() &&
		rightData.IsBalance() &&
		abs(leftData.Height(), rightData.Height()) < 2

	return NewNodeBalanceData(height, isBalance)
}

// 完全二叉树的判断
func IsCBT(root *BinaryTree) bool {
	return true
}

// 宽度优先遍历（队列）
// 1. 先进头节点
// 2. 队列不为空，就弹出一个节点，再把弹出节点的左右节点依次入队（先左后右）,重复步骤2，知道队列为空
func BreadthFirstTraversal(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	queue := NewQueue()
	queue.Enqueue(root)
	for !queue.Empty() {
		node := queue.Dequeue().(*BinaryTree)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
		vals = append(vals, node.Value)
	}
	return
}

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
