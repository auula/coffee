package list

import "github.com/auula/coffee"

type Node[V any] struct {
	Perv, Next *Node[V]
	Value      V
}

type List[V any] struct {
	Head, Tail *Node[V]
	Size       uint
}

func New[V any]() List[V] {
	return List[V]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func NewNode[V any](v V) *Node[V] {
	return &Node[V]{Value: v}
}

func (list *List[V]) RPush(v V) {
	node := NewNode(v)

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Perv = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}

	list.Size += 1
}

func (list *List[V]) LPush(v V) {
	node := NewNode(v)

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Perv = node
		list.Head = node
	}

	list.Size += 1
}

func (list *List[V]) Front() V {
	node := list.Head
	list.Head = node.Next
	return node.Value
}

func (list *List[V]) Back() V {
	node := list.Tail
	list.Tail = node.Perv
	return node.Value
}

func (list *List[V]) Get(index uint) *Node[V] {

	if index > list.Size {
		return nil
	}

	if list.Size == 0 {
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
	return list.Size != 0
}

func (list *List[V]) Next() V {
	value := list.Front()
	list.Size -= 1
	return value
}
