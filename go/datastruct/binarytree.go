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
func PreOrderByRecursive(root *BinaryTree) (vals []int) {
	var preorder func(*BinaryTree)
	preorder = func(bt *BinaryTree) {
		if bt == nil {
			return
		}
		vals = append(vals, bt.Value)
		preorder(bt.Left)
		preorder(bt.Right)
	}
	preorder(root)
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
func InOrderByRecursive2(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}

	vleft := InOrderByRecursive2(root.Left)
	vals = append(vals, vleft...)
	// 处理当前节点
	vals = append(vals, root.Value)
	vright := InOrderByRecursive2(root.Right)
	vals = append(vals, vright...)
	return
}

// 中序遍历（递归）
func InOrderByRecursive(root *BinaryTree) (vals []int) {
	var inoder func(*BinaryTree)
	inoder = func(bt *BinaryTree) {
		if bt == nil {
			return
		}
		inoder(root.Left)
		vals = append(vals, bt.Value)
		inoder(root.Right)
	}
	inoder(root)
	return
}

// 中序遍历（非递归）
// 1. 申请一个栈stack，记住cur=root
// 2. 先把 cur 压入栈，cur=cur.left 重复2,找到最左边的一个节点
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
	var postOrder func(*BinaryTree)
	postOrder = func(bt *BinaryTree) {
		if bt == nil {
			return
		}
		postOrder(bt.Left)
		postOrder(bt.Right)
		vals = append(vals, bt.Value)
	}
	postOrder(root)
	return
}

func PostOrderByRecursive2(root *BinaryTree) (vals []int) {
	if root == nil {
		return
	}
	vleft := PostOrderByRecursive2(root.Left)
	vright := PostOrderByRecursive2(root.Right)
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
		max = Max(max, leftData.Max())
		min = Min(min, leftData.Min())
	}
	if node.Right != nil {
		max = Max(max, rightData.Max())
		min = Min(min, rightData.Min())
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

	height := Max(leftData.height, rightData.height) + 1
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

	height := Max(leftData.Height(), rightData.Height()) + 1
	isBalance := leftData.IsBalance() &&
		rightData.IsBalance() &&
		Abs(leftData.Height()-rightData.Height()) < 2

	return NewNodeBalanceData(height, isBalance)
}

// 完全二叉树的判断
// 使用宽度优先遍历
func IsCBT(root *BinaryTree) bool {
	if root == nil {
		return true
	}
	// 某一个节点左右孩子不双全，遇到就标记为 true
	leaf := false
	queue := NewQueue()
	queue.Enqueue(root)
	for !queue.Empty() {
		node := queue.Dequeue().(*BinaryTree)
		l := node.Left
		r := node.Right

		// 1. 当前节点 node 左右孩子不双全,并且不是叶子节点
		// 2. 或者左孩子为空，右孩子不为空
		// 可以判定该节点以下的子树不为完全二叉树
		if (leaf && (l != nil || r != nil)) ||
			(l == nil && r != nil) {
			return false
		}
		if l != nil {
			queue.Enqueue(l)
		}
		if r != nil {
			queue.Enqueue(r)
		}
		if l == nil || r == nil {
			leaf = true
		}
	}
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

// 给点 node1,node2, 找到他们最低公共祖先节点
// node1,node2 一定属于 head 为头的树
func LowestAncestor1(head, node1, node2 *BinaryTree) *BinaryTree {
	if head == nil || node1 == head || node2 == head {
		return head
	}
	fatherMap := map[*BinaryTree]*BinaryTree{
		head: head,
	}
	nodeFatherMap(head, fatherMap)
	// set1 放置 node1 的父节点
	set1 := make(map[*BinaryTree]*BinaryTree)
	cur := node1
	for cur != fatherMap[cur] {
		set1[cur] = cur
		cur = fatherMap[cur]
	}
	cur = node2
	// 查看 node2 的父节点是否存在于 set1 中
	for cur != fatherMap[cur] {
		if node, ok := set1[cur]; ok {
			return node
		}
		cur = fatherMap[cur]
	}
	return nil
}

// 给点 node1,node2, 找到他们最低公共祖先节点
// node1,node2 一定属于 head 为头的树
func LowestAncestor2(head, node1, node2 *BinaryTree) *BinaryTree {
	if head == nil || node1 == head || node2 == head {
		return head
	}

	left := LowestAncestor2(head.Left, node1, node2)
	right := LowestAncestor2(head.Right, node1, node2)
	if left != nil && right != nil { // 左右子树都不为空，说明当前节点 head 为他们最小的公共祖先
		return head
	}
	if left != nil { // 如果左不为空，那结果就是左边返回的信息，右边亦是如此
		return left
	}
	return right
}

// 给点 node1,node2, 找到他们最低公共祖先节点
// node1,node2 一定属于 head 为头的树
// 并查集
func LowestAncestor3(head, node1, node2 *BinaryTree) *BinaryTree {
	if head == nil || node1 == head || node2 == head {
		return head
	}
	return nil
}

// 构建节点的父节点关系
func nodeFatherMap(head *BinaryTree, fatherMap map[*BinaryTree]*BinaryTree) {
	if head == nil {
		return
	}
	fatherMap[head.Left] = head
	fatherMap[head.Right] = head
	nodeFatherMap(head.Left, fatherMap)
	nodeFatherMap(head.Right, fatherMap)
}

type BTNode struct {
	Value               int
	Left, Right, Parent *BTNode
}

// 后继节点问题（后续节点是中序遍历的下一个）
func GetSuccessorNode(root *BTNode) *BTNode {
	if root == nil {
		return root
	}
	if root.Right != nil {
		return getLastLeft(root.Right)
	}
	parent := root.Parent
	for parent != nil && parent.Left != root {
		root = parent
		parent = root.Parent
	}
	return parent
}

func getLastLeft(node *BTNode) *BTNode {
	if node == nil {
		return node
	}
	for node != nil {
		node = node.Left
	}
	return node
}

// 凹凸问题
func PrintDetail(N int) {
	printProcess(1, N, true)
}

// 递归到某一个节点
// i 为节点层数，N 表示总层数，down==true 凹 down==false 凸
func printProcess(i, N int, down bool) {
	if i > N {
		return
	}
	printProcess(i+1, N, true)
	var res string
	if down {
		res = "凹"
	} else {
		res = "凸"
	}
	fmt.Printf("%s\t", res)
	printProcess(i+1, N, false)
}
