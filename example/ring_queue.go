package main

import (
	"fmt"

	"github.com/auula/coffee"
	"github.com/auula/coffee/ring"
)

func main() {

	rq := ring.New[int](10)

	for i := 0; i < 10; i++ {
		rq.Put(i)
	}

	fmt.Println("Size:", rq.Size())
	fmt.Println("IsFull:", rq.IsFull())

	// rq.Get()
	coffee.ForEach(rq.Iter(), func(i int) {
		fmt.Println(i)
	})

	fmt.Println("Size:", rq.Size())
	fmt.Println("IsFull:", rq.IsFull())
}
