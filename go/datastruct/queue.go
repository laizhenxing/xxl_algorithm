package datastruct

import (
	"container/list"
)

type Queue struct {
	queue *list.List
}

func NewQueue() *Queue {
	return &Queue{
		queue: list.New(),
	}
}

// 入队
func (q *Queue) Enqueue(val interface{}) {
	q.queue.PushFront(val)
}

// 出队
func (q *Queue) Dequeue() interface{} {
	if q.queue.Len() > 0 {
		elem := q.queue.Back()
		q.queue.Remove(elem)
		return elem.Value
	}
	return nil
}

// 查看头元素
func (q *Queue) Peek() interface{} {
	elem := q.queue.Front()
	if elem == nil {
		return nil
	}
	return elem.Value
}

// 获取队列长度
func (q *Queue) Len() int {
	return q.queue.Len()
}

// 队列是否为空
func (q *Queue) Empty() bool {
	return q.queue.Len() == 0
}

/***********************************范型队列**********************************************/
type QueueL[T interface{}] struct {
	elements []T
}

// 输入进入队列尾部
func (q *QueueL[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// 从队列头部取出并删除对应数据
func (q *QueueL[T]) Dequeue() (T, bool) {
	var value T
	if q.Len() == 0 {
		return value, true
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, q.Len() == 0
}

// 队列长度
func (q *QueueL[T]) Len() int {
	return len(q.elements)
}
