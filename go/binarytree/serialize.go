package binarytree

import (
	"container/list"
	"strconv"
	"strings"
)

type ISerialize interface {
	Serialize(root *TreeNode) string
	Deserialize(data string) *TreeNode
}

var (
	SEP  string = ","
	NULL string = "#"
)

type Codec struct{}

func (c *Codec) Serialize(root *TreeNode) string {
	var sb strings.Builder
	c.serialize(root, &sb)
	return sb.String()
}

func (c *Codec) serialize(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(NULL)
		sb.WriteString(SEP)
		return
	}

	// 前序遍历
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)

	c.serialize(root.Left, sb)
	c.serialize(root.Right, sb)
}

func (c *Codec) Deserialize(data string) *TreeNode {
	sdata := strings.Split(data, SEP)

	l := list.New()
	for _, v := range sdata {
		l.PushBack(v)
	}
	return c.deserialize(l)
	// return c.deserialize_with_slice(&sdata)
}

func (c *Codec) deserialize(data *list.List) *TreeNode {
	if data.Len() < 1 {
		return nil
	}

	// 移除第一个元素，它是根节点
	first := data.Remove(data.Front())

	if first == NULL {
		return nil
	}

	val, _ := strconv.Atoi(first.(string))
	root := &TreeNode{Val: val}

	root.Left = c.deserialize(data)
	root.Right = c.deserialize(data)

	return root
}

func (c *Codec) deserialize_with_slice(data *[]string) *TreeNode {
	if len(*data) < 1 {
		return nil
	}
	// 移除第一个元素，它是根节点
	first := (*data)[0]
	*data = (*data)[1:]
	if first == NULL {
		return nil
	}

	val, _ := strconv.Atoi(first)
	root := &TreeNode{Val: val}

	root.Left = c.deserialize_with_slice(data)
	root.Right = c.deserialize_with_slice(data)

	return root
}

func Serialize(root *TreeNode) string {
	return new(Codec).Serialize(root)
}

func Deserialize(data string) *TreeNode {
	return new(Codec).Deserialize(data)
}
