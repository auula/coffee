package list

import (
	"sync"

	"github.com/auula/coffee"
)

type Node[V any] struct {
	Prev, Next *Node[V]
	Value      V
}

type List[V any] struct {
	Head, Tail *Node[V]
	size       int
	sync.Mutex
}

func New[V any]() List[V] {
	return List[V]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func NewNode[V any](v V) *Node[V] {
	return &Node[V]{Value: v}
}

func (list *List[V]) RPush(v V) {
	node := NewNode(v)

	if list.size == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Prev = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}

	list.size += 1
}

func (list *List[V]) LPush(v V) {
	node := NewNode(v)

	if list.size == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Prev = node
		list.Head = node
	}

	list.size += 1
}

func (list *List[V]) Front() *Node[V] {
	node := list.Head
	list.Head = node.Next
	return node
}

func (list *List[V]) Back() *Node[V] {
	node := list.Tail
	list.Tail = node.Prev
	return node
}

func (list *List[V]) Get(index int) *Node[V] {

	if index > list.size {
		return nil
	}

	if list.size == 0 {
		return nil
	}

	node := list.Head
	for i := 0; i < int(index); i++ {
		node = node.Next
	}

	return node
}

func (list *List[V]) Iter() coffee.Iterator[V] {
	return list
}

func (list *List[V]) HasNext() bool {
	return list.Head != nil && list.Tail != nil
}

func (list *List[V]) Next() V {
	return list.Front().Value
}

func (list *List[V]) Size() int {
	return list.size
}
