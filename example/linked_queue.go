//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/auula/coffee/queue"
)

func main() {

	var q queue.Queued[int] = queue.New[int]()

	for i := 0; i < 10; i++ {
		q.EnQueue(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(*q.DeQueue())
	}
}
