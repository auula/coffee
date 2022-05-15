package tree

import "github.com/auula/coffee"

type Node[T coffee.Number] struct {
	Data                  T
	LeftChild, RightChild *Node[T]
}

type BST[T coffee.Number] struct {
	Root *Node[T]
	Size int
}

func (bt *BST[T]) Insert(v T) {

	if bt.Root == nil {
		bt.Root = &Node[T]{Data: v}
		return
	}

	node := bt.Root

	for node != nil {
		if v > node.Data {
			if node.RightChild == nil {
				node.RightChild = &Node[T]{Data: v}
				return
			}
			node = node.RightChild
		} else {
			if node.LeftChild == nil {
				node.LeftChild = &Node[T]{Data: v}
				return
			}
			node = node.LeftChild
		}
	}
}

func (bt *BST[T]) Search(v T) *Node[T] {
	node := bt.Root
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
