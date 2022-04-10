package queue

import "sync"

type Node[V any] struct {
	Perv, Next *Node[V]
	Value      V
}

type LinkedQueue[V any] struct {
	head, tail *Node[V]
	size       int
	sync.Mutex
}

func (q *LinkedQueue[V]) EnQueue(v V) bool {
	q.Lock()
	defer q.Unlock()
	node := NewNode(v)
	if q.size == 0 {
		q.head = node
		q.tail = node
	}
	node.Perv = q.tail
	q.tail.Next = node
	q.tail = node
	q.size += 1
	return true
}

func (q *LinkedQueue[V]) DeQueue() *V {
	q.Lock()
	defer q.Unlock()
	if q.size == 0 {
		return nil
	}
	node := q.head
	q.head = node.Next
	q.size--
	return &node.Value
}

func (q *LinkedQueue[V]) IsFull() bool {
	return false
}

func (q *LinkedQueue[V]) Size() int {
	q.Lock()
	defer q.Unlock()
	return q.size
}

func NewNode[V any](v V) *Node[V] {
	return &Node[V]{Value: v}
}

func New[V any]() Queued[V] {
	return new(LinkedQueue[V])
}