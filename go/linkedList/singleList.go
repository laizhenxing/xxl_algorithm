package linkedlist

type SingleList struct {
	Val  int
	Next *SingleList
}

func NewSingleList(val int) *SingleList {
	return &SingleList{
		Val:  val,
		Next: nil,
	}
}

// func BuildSingleList(vals []int) *SingleList {
// 	var head, cur *SingleList
// 	for _, v := range vals {
// 		cur = &SingleList{
// 			Val: v,
// 		}
// 	}

// 	return head
// }

// 翻转列表
func Reverse(head *SingleList) *SingleList {
	if head == nil {
		return nil
	}

	last := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var successor *SingleList

// 翻转前n个节点
func ReversN(head *SingleList, n int) *SingleList {
	if n == 1 {
		successor = head.Next
		return head
	}

	last := ReversN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor

	return last
}

// 翻转 [m,n]个节点
func ReverseBetween(head *SingleList, m, n int) *SingleList {
	if m == 1 {
		return ReversN(head, n)
	}
	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}
