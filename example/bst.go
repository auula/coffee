package main

import (
	"fmt"

	"github.com/auula/coffee/bst"
)

func main() {
	tree := bst.New[int]()
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(16)
	tree.Insert(3)
	tree.Insert(10)

	tree.Delete(9)

	channel := make(chan *bst.Node[int])

	go func() {
		bst.Previous(tree.Root, channel)
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

}
