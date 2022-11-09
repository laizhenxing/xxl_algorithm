package datastruct

import (
	"container/heap"
	"fmt"
)

// 堆: 特殊的完全二叉树
// 数组堆化
// 父节点：(i-1)/2
// 左节点：2*i+1
// 右节点：2*i+2
// 大顶堆，小顶堆

// 优先级队列就是堆
type Heap struct {
	heap.Interface
}

// 堆插入
func HeapInsert(arr []int, value, index int) {
	arr[index] = value
	for index != 0 {
		parent := (index - 1) >> 1
		if arr[index] > arr[parent] {
			Swap(arr, index, parent)
			index = parent
		} else {
			break
		}
	}
}

// 堆化
func Heapify(arr []int, index, heapSize int) {
	left, right, largest := index<<1+1, index<<1+2, index
	for left < heapSize {
		if arr[left] > arr[index] {
			largest = left
		}
		if right < heapSize && arr[right] > arr[largest] {
			largest = right
		}
		if largest != index {
			Swap(arr, largest, index)
		} else {
			break
		}
		index = largest
		left = index<<1 + 1
		right = index<<1 + 2
	}
}

// 堆排序 O(N*logN)
func HeapSort(arr []int) {
	arrLen := len(arr)
	if arr == nil || arrLen < 2 {
		return
	}
	// for i := 0; i < len(arr); i++ { // O(N)
	// 	HeapInsert(arr, arr[i], i) // O(logN)
	// }
	for i := arrLen - 1; i >= 0; i-- {
		Heapify(arr, i, arrLen)
	}
	heapSize := arrLen
	heapSize--
	Swap(arr, 0, heapSize)
	for heapSize > 0 { // O(N)
		Heapify(arr, 0, heapSize) // O(logN)
		heapSize--
		Swap(arr, 0, heapSize)
	}
}

// 实现 heap.Interface，就可以使用内置的堆的操作
// 内置的堆就是优先级队列
type PriorityQueue []int

// Len is the number of elements in the collection.
func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i int, j int) bool {
	return q[i] < q[j]
}

// Swap swaps the elements with indexes i and j.
func (q PriorityQueue) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	v := old[n-1]
	*q = old[0 : n-1]
	return v
}

func (q *PriorityQueue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func PQ() {
	h := &PriorityQueue{3, 2, 1, 23, 7, 4, 0}
	heap.Init(h)
	heap.Push(h, 90)
	for h.Len() > 0 {
		fmt.Println("min: ", heap.Pop(h))
	}
}
