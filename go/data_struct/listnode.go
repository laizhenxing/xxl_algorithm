package datastruct

import (
	"fmt"
	"sort"
)

// 反转单向和双向链表
// 要求：链表长度为N，时间复杂度要求O(N)，空间复杂度 O(1)

type SingleNode struct {
	Value int
	Next  *SingleNode
}

func (sn *SingleNode) GetDetails() (length int, vals []int) {
	head := sn.Next
	for {
		if head == nil {
			break
		}
		length++
		vals = append(vals, head.Value)
		head = head.Next
	}

	return
}

func (sn *SingleNode) Display() {
	if sn == nil {
		fmt.Println("链表为空")
		return
	}

	tmp := sn
	for {
		if tmp == nil {
			break
		}
		fmt.Println(tmp)
		tmp = tmp.Next
	}

}

func (sn *SingleNode) Add(val int) {
	node := &SingleNode{
		Value: val,
		Next:  nil,
	}

	cur := sn
	for {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}

	cur.Next = node
}

// 添加多个节点到链表
func (sn *SingleNode) MulAdd(vals []int) {
	sort.Ints(vals)
	for _, val := range vals {
		sn.Add(val)
	}
}

func CreateLinkWithVals(vals []int) (sn *SingleNode) {
	sn = &SingleNode{}
	for _, val := range vals {
		sn.Add(val)
	}
	return
}

// 反转单向链表
func ReverseList(head *SingleNode) *SingleNode {
	var pre, next *SingleNode = nil, nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

type DoblyNode struct {
	Value int
	Pre   *DoblyNode
	Next  *DoblyNode
}

func ReverseListWithHead(head *SingleNode) *SingleNode {
	oldHead := head
	head = ReverseList(head.Next)
	oldHead.Next = head
	return oldHead
}

// 反转双向链表
func ReverseDoblyList(head *DoblyNode) *DoblyNode {
	var pre, next *DoblyNode = new(DoblyNode), new(DoblyNode)
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Pre = next
		pre = head
		head = next
	}
	return pre
}

// 打印两个有序链表的公共部分
// 要求：空间复杂度 O(1)
func SingleNodeFindCommPart(head1, head2 *SingleNode) (comm []int) {
	n1 := head1.Next
	n2 := head2.Next
	for {
		if n1 == nil || n2 == nil { // 	某一个链表移动完就结束
			break
		}
		if n1.Value < n2.Value { // n1 小，n1 移动
			n1 = n1.Next
		} else if n2.Value < n1.Value { // n2 小，n2 移动
			n2 = n2.Next
		} else { // 相同就是共有的值，一起移动
			comm = append(comm, n1.Value)
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	return
}

// 链表重合第一个节点

// 题目：判断回文
// 1. 使用栈
func IsPalindrome1(node *SingleNode) bool {
	head := node.Next
	if head == nil || head.Next == nil {
		return true
	}

	// 模拟栈
	vals := make([]int, 0)
	// 将链表元素入栈
	for {
		if head == nil {
			break
		}
		vals = append(vals, head.Value)
		head = head.Next
	}
	head = node.Next
	vlen := len(vals)
	i := 1
	for {
		if head == nil {
			break
		}
		if head.Value != vals[vlen-i] {
			return false
		}
		head = head.Next
		i++
	}
	return true
}

// 判断回文
// 2. 右边部分入栈，使用快慢指针，快慢指针可以找中点
func IsPalindrome2(node *SingleNode) bool {
	if node.Next == nil || node == nil {
		return true
	}
	n1, n2 := node.Next, node.Next
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next      // n1 --> 中部
		n2 = n2.Next.Next // m2 --> 尾部
	}
	n2 = n1.Next
	vals := make([]int, 0)
	for n2 != nil { // 右边入栈
		vals = append(vals, n2.Value)
		n2 = n2.Next
	}
	n1 = node.Next
	for i, vlen := 0, len(vals); i < vlen; i++ {
		// 弹出栈，对比左边的值是否对称
		if n1.Value != vals[vlen-1-i] {
			return false
		}
		n1 = n1.Next
	}

	return true
}

// 3. 反转右边部分链表，使用3个变量
func IsPalindrome3(node *SingleNode) bool {
	if node.Next == nil || node == nil {
		return true
	}

	n1, n2 := node.Next, node.Next
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next      // n1 --> 中部
		n2 = n2.Next.Next // m2 --> 尾部
	}
	n2 = n1.Next

	n1 = nil
	var n3 *SingleNode
	for n2 != nil { // 反转右边链表
		n3 = n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}

	n2 = n1 // n1 为左边最后一个节点
	n3 = n1 // n3 保存最后一个节点

	// n2 为右边第一个点
	res := true
	n1 = node.Next  // n1 为左边第一个点
	for n2 != nil { // 判断右边部分是否与左边部分对称
		if n1.Value != n2.Value {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	// 恢复链表【反转右边部分链表】
	n1 = nil // n1 == pre | n3 == cur | n2 == next
	for n3 != nil {
		n2 = n3.Next
		n3.Next = n1
		n1 = n3
		n3 = n2
	}

	return res
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	s := head // slow
	f := head // fast
	for f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	// s 此时在中点位置，反转右侧
	s = s.Next
	cur := s
	s = nil
	for cur != nil { // 反转右侧链表
		f = cur.Next
		cur.Next = s
		s = cur
		cur = f
	}

	res := true
	f = head
	cur = s // 保存最后一个节点位置
	for s != nil {
		if s.Val != f.Val {
			res = false
			break
		}
		s = s.Next
		f = f.Next
	}

	// mid := f
	// 恢复右侧链表
	// s = nil
	for cur != nil {
		f = cur.Next
		cur.Next = s
		s = cur
		cur = f
	}
	// mid.Next = s

	return res
}
