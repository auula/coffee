package queue

import (
	"github.com/auula/coffee"
)

type Node[V any] struct {
	Perv, Next *Node[V]
	Value      V
}

type LinkedQueue[V any] struct {
	head, tail *Node[V]
	size       int
}

func (q *LinkedQueue[V]) EnQueue(v V) bool {
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
	return q.size
}

func NewNode[V any](v V) *Node[V] {
	return &Node[V]{Value: v}
}

func New[V any]() Queued[V] {
	return new(LinkedQueue[V])
}

func (q *LinkedQueue[V]) Iter() coffee.Iterator[V] {
	return q
}

func (q *LinkedQueue[V]) HasNext() bool {
	return q.size != 0
}

func (q *LinkedQueue[V]) Next() V {
	return *q.DeQueue()
}
