// Open Source: MIT License
// Author: Leon Ding <ding@ibyte.me>
// Date: 2022/4/10 - 7:09 PM - UTC/GMT+08:00

package main

import (
	"fmt"
	"github.com/auula/coffee/queue"
)

func main() {
	var q queue.Queuer[int]

	q = queue.NewArray[int](10)

	for i := 10; i > 0; i-- {
		q.EnQueue(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(q.DeQueue())
	}

	fmt.Println("IsFull:", q.IsFull())

	fmt.Println("Size:", q.Size())
}
