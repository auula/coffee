package bst

import "github.com/auula/coffee"

type Node[T coffee.Number] struct {
	Data                  T
	LeftChild, RightChild *Node[T]
}

type Tree[T coffee.Number] struct {
	root *Node[T]
	size int
}

func New[T coffee.Number]() *Tree[T] {
	return new(Tree[T])
}

func (bt *Tree[T]) Insert(v T) {

	if bt.root == nil {
		bt.root = &Node[T]{Data: v}
		bt.size += 1
		return
	}

	node := bt.root

	for node != nil {
		if v > node.Data {
			if node.RightChild == nil {
				node.RightChild = &Node[T]{Data: v}
				bt.size += 1
				return
			}
			node = node.RightChild
		} else {
			if node.LeftChild == nil {
				node.LeftChild = &Node[T]{Data: v}
				bt.size += 1
				return
			}
			node = node.LeftChild
		}
	}
}

func (bt *Tree[T]) Search(v T) *Node[T] {
	node := bt.root
	for node != nil {
		if v < node.Data {
			node = node.LeftChild
		} else if v > node.Data {
			node = node.RightChild
		} else {
			return node
		}
	}
	return nil
}

func Previous[T coffee.Number](node *Node[T], channel chan *Node[T]) {
	if node == nil {
		return
	}
	channel <- node
	Previous(node.LeftChild, channel)
	Previous(node.RightChild, channel)
}

func Intermediate[T coffee.Number](node *Node[T], channel chan *Node[T]) {
	if node == nil {
		return
	}
	Intermediate(node.LeftChild, channel)
	channel <- node
	Intermediate(node.RightChild, channel)
}

func Post[T coffee.Number](node *Node[T], channel chan *Node[T]) {
	if node != nil {
		return
	}
	Post(node.LeftChild, channel)
	Post(node.RightChild, channel)
	channel <- node
}
