package dp

import (
	"algo/datastruct"
)

// 并查集
// 链表实现
type DisjoinSetNode struct {
	depth    int // 子节点的数量
	parent   *DisjoinSetNode
	children *datastruct.Set
	Data     interface{}
}

func NewDisjoinSetNode() *DisjoinSetNode {
	ds := &DisjoinSetNode{
		children: datastruct.NewSet(),
	}
	ds.parent = ds
	return ds
}

func (ds *DisjoinSetNode) FindHead() *DisjoinSetNode {
	childrenStack := datastruct.NewStack()
	for ds != ds.parent {
		childrenStack.Push(ds)
		ds.parent = ds.parent.parent
		ds = ds.parent
	}
	for !childrenStack.IsEmpty() {
		child := childrenStack.Pop().(*DisjoinSetNode)
		child.parent = ds
		ds.children.Store(child)
	}
	return ds
}

func (ds *DisjoinSetNode) IsSameSet(d1 *DisjoinSetNode) bool {
	return ds.FindHead() == d1.FindHead()
}

func (ds *DisjoinSetNode) Union(d1, d2 *DisjoinSetNode) {
	d1Root := d1.FindHead()
	d2Root := d2.FindHead()
	if d1Root != d2Root {
		d1Root.depth += d2Root.depth
		switch {
		case d1Root.depth > d2Root.depth:
		case d1Root.depth == d2Root.depth:
			children := d2Root.children.Clear()
			for _, child := range children {
				node := child.(*DisjoinSetNode)
				d1Root.children.Store(node)
				node.parent = d1Root
			}
			d2Root.parent = d1Root
		case d1Root.depth < d2Root.depth:
			children := d1Root.children.Clear()
			for _, child := range children {
				node := child.(*DisjoinSetNode)
				d2Root.children.Store(node)
				node.parent = d2Root
			}
			d1Root.parent = d2Root
		}
	}
}
