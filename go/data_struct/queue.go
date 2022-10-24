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
