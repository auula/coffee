package tree

type Node[T comparable] struct {
	Data                  T
	LeftChild, RightChild *Node[T]
}

type BinaryTree[T comparable] struct {
	Root *Node[T]
	Size int
}

func (bt *BinaryTree[T]) Insert(v T) {

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
