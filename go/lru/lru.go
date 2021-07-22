package lru

import (
	"sync"
)

/**********/
/*  Node  */
/**********/
type Node struct {
	key, val  interface{}
	Pre, Next *Node
}

func NewNode(k, v interface{}) *Node {
	return &Node{
		key: k,
		val: v,
	}
}

func (n *Node) GetKey() interface{} {
	return n.key
}

func (n *Node) GetVal() interface{} {
	return n.val
}

func (n *Node) Tomap() map[interface{}]interface{} {
	node := make(map[interface{}]interface{})
	node[n.key] = n.val
	return node
}

/****************/
/*  DoubleList  */
/****************/

type DoubleList struct {
	head, tail *Node
	size       int
}

func NewDoubleList() *DoubleList {
	dl := &DoubleList{
		head: NewNode(0, 0),
		tail: NewNode(0, 0),
		size: 0,
	}
	dl.head.Next = dl.tail
	dl.tail.Pre = dl.head
	return dl
}

func (d *DoubleList) AddLast(node *Node) {
	node.Next = d.tail.Pre
	node.Pre = d.tail
	d.tail.Pre.Next = node
	d.tail.Pre = node
	d.size++
}

func (d *DoubleList) Remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	d.size--
}

func (d *DoubleList) RemoveFirst() *Node {
	if d.head == d.tail {
		return nil
	}
	first := d.head.Next
	d.Remove(first)
	return first
}

func (d *DoubleList) First() map[interface{}]interface{} {
	if d.head == d.tail {
		return nil
	}
	node := d.tail.Pre
	return node.Tomap()
}

func (d *DoubleList) Size() int {
	return d.size
}

/**************/
/*  LRUCache  */
/**************/
type LRUCache struct {
	keyMap *sync.Map
	cache  *DoubleList
	cap    int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		keyMap: new(sync.Map),
		cache:  NewDoubleList(),
		cap:    capacity,
	}
}

// 将某个 key 提升为最近使用
func (lru *LRUCache) makeRecently(key interface{}) {
	if node := lru.getNode(key); node != nil {
		lru.cache.Remove(node)
		lru.cache.AddLast(node)
	}
}

// 添加最近使用元素
func (lru *LRUCache) addRecently(key, val interface{}) {
	node := NewNode(key, val)
	lru.keyMap.Store(key, node)
	lru.cache.AddLast(node)
}

// 删除某个key
func (lru *LRUCache) deleteKey(key interface{}) {
	if node := lru.getNode(key); node != nil {
		lru.cache.Remove(node)
		lru.keyMap.Delete(key)
	}
}

// 删除最久没有使用
func (lru *LRUCache) removeLeastRecently() {
	node := lru.cache.RemoveFirst()
	lru.keyMap.Delete(node.GetKey())
}

func (lru *LRUCache) getNode(key interface{}) *Node {
	if inode, ok := lru.keyMap.Load(key); ok {
		if node, ok := inode.(*Node); ok {
			return node
		}
	}
	return nil
}

func (lru *LRUCache) Get(key interface{}) interface{} {
	if node := lru.getNode(key); node != nil {
		lru.makeRecently(key)
		return node.GetVal()
	}
	return nil
}

func (lru *LRUCache) Put(key, val interface{}) {
	if node := lru.getNode(key); node != nil {
		lru.deleteKey(key)
		lru.addRecently(key, val)
		return
	}

	if lru.cap == lru.cache.Size() {
		lru.removeLeastRecently()
	}

	lru.addRecently(key, val)
}

func (lru *LRUCache) First() map[interface{}]interface{} {
	return lru.cache.First()
}
