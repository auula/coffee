//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/auula/coffee/queue"
)

func main() {
	var q queue.Queued[int]

	q = queue.NewArray[int](10)

	for i := 10; i > 0; i-- {
		q.EnQueue(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(*q.DeQueue())
	}

	fmt.Println("IsFull:", q.IsFull())

	fmt.Println("Size:", q.Size())
}
